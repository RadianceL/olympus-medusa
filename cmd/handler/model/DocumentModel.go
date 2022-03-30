package model

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/data"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
	"medusa-globalization-copywriting-system/cmd/handler/model/basic"
	"medusa-globalization-copywriting-system/tools/language"
	"time"
)

const (
	// ApplicationModelTableName tb_application
	documentTableName          = "tb_application_globalization_document_code"
	applicationIdField         = "application_id"
	namespaceIdField           = "namespace_id"
	documentCodeField          = "document_code"
	isEnableField              = "is_enable"
	onlineTimeField            = "online_time"
	onlineOperatorUserIdField  = "online_operator_user_id"
	offlineTimeField           = "offline_time"
	offlineOperatorUserIdField = "offline_operator_user_id"
	offlineAccessUserIdField   = "offline_access_user_id"
	createTimeField            = "create_time"
	createUserIdField          = "create_user_id"
	deleteFlagField            = "delete_flag"
	deleteTimeField            = "delete_time"
	deleteUserIdField          = "delete_user_id"

	documentValueTableName  = "tb_application_globalization_document_value"
	documentIdField         = "document_id"
	countryIsoField         = "country_iso"
	countryNameField        = "country_name"
	documentValueField      = "document_value"
	documentIsOnlineField   = "document_is_online"
	updateTimeField         = "update_time"
	updateUserIdField       = "update_user_id"
	lastUpdateDocumentField = "last_update_document"
)

// DocumentModel is application model structure.
type DocumentModel struct {
	basic.BaseModel
}

func init() {
	RepositoryModelContainer.Register(Document())
}

func (documentModel DocumentModel) Initialization() {
	DocumentHandler = Document().SetConn(datasource.Conn)
	println("初始化ApplicationHandler")
}

// Document return a default application model.
func Document() DocumentModel {
	return DocumentModel{BaseModel: basic.BaseModel{TableName: documentTableName}}
}

func (documentModel DocumentModel) SetConn(connection datasource.Connection) DocumentModel {
	documentModel.Conn = connection
	return documentModel
}

func (documentModel DocumentModel) CreateDocument(namespaceRequest *Entity.GlobalDocumentRequest) (int64, error) {
	tx := documentModel.Conn.BeginTx()
	insertDocumentCodeResult, err := documentModel.Table(documentTableName).
		WithTx(tx).
		Insert(dialect.H{
			applicationIdField: namespaceRequest.ApplicationId,
			namespaceIdField:   namespaceRequest.NamespaceId,
			onlineTimeField:    time.Now(),
			documentCodeField:  namespaceRequest.Key,
			createUserIdField:  0,
		})
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	documents := namespaceRequest.Documents
	for _, document := range documents {
		languageCountry := language.FindLanguage(document.CountryCode)
		if languageCountry == nil {
			return 0, errors.New("未识别的国家编码，请检查后重试")
		}
		_, err := documentModel.Table(documentValueTableName).
			WithTx(tx).
			Insert(dialect.H{
				documentIdField:    insertDocumentCodeResult,
				namespaceIdField:   namespaceRequest.NamespaceId,
				countryIsoField:    document.CountryCode,
				countryNameField:   languageCountry.CountryName,
				documentValueField: document.LanguageValue,
			})
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}
	commitError := tx.Commit()
	if commitError != nil {
		_ = tx.Rollback()
		return 0, commitError
	}
	return insertDocumentCodeResult, err
}

func (documentModel DocumentModel) SearchDocumentByNamespaceId(applicationId int, namespaceId int) ([]data.TableGlobalDocument, error) {
	statement := documentModel.Table(documentTableName).Select("*")
	statement.Where(namespaceIdField, "=", namespaceId)
	statement.Where(applicationIdField, "=", applicationId)
	documentMaps, err := statement.All()
	if err != nil {
		return []data.TableGlobalDocument{}, err
	}
	if len(documentMaps) <= 0 {
		return []data.TableGlobalDocument{}, nil
	}

	var resultList []data.TableGlobalDocument
	for _, document := range documentMaps {
		var documentResult data.TableGlobalDocument
		mapstructure.Decode(document, &documentResult)

		queryDocumentValueStatement := documentModel.Table(documentValueTableName).Select("*")
		queryDocumentValueStatement.Where(documentIdField, "=", documentResult.Id)
		documentValueResultDataMaps, documentValueErr := queryDocumentValueStatement.All()
		if documentValueErr != nil {
			return []data.TableGlobalDocument{}, documentValueErr
		}
		var result []data.TableGlobalDocumentLanguage
		for _, documentValueResultData := range documentValueResultDataMaps {
			var tableGlobalDocumentLanguageOutputResult data.TableGlobalDocumentLanguage
			mapstructure.Decode(documentValueResultData, &tableGlobalDocumentLanguageOutputResult)
			result = append(result, tableGlobalDocumentLanguageOutputResult)
		}
		documentResult.Documents = result
		resultList = append(resultList, documentResult)
	}
	return resultList, nil
}

func (documentModel DocumentModel) SearchDocumentById(documentId int) (data.TableGlobalDocument, error) {
	statement := documentModel.Table(documentTableName).Select("*")
	statement.Where(id, "=", documentId)
	resultData, err := statement.All()
	if err != nil {
		return data.TableGlobalDocument{}, err
	}
	if len(resultData) <= 0 {
		return data.TableGlobalDocument{}, errors.New("未查询到编码信息，请确认后重试")
	}
	var outputResult data.TableGlobalDocument
	mapstructure.Decode(resultData[0], &outputResult)

	queryDocumentValueStatement := documentModel.Table(documentValueTableName).Select("*")
	queryDocumentValueStatement.Where(documentIdField, "=", documentId)
	documentValueResultDataMaps, documentValueErr := queryDocumentValueStatement.All()
	if documentValueErr != nil {
		return data.TableGlobalDocument{}, documentValueErr
	}
	var result []data.TableGlobalDocumentLanguage
	for _, documentValueResultData := range documentValueResultDataMaps {
		var tableGlobalDocumentLanguageOutputResult data.TableGlobalDocumentLanguage
		mapstructure.Decode(documentValueResultData, &tableGlobalDocumentLanguageOutputResult)
		result = append(result, tableGlobalDocumentLanguageOutputResult)
	}
	outputResult.Documents = result
	return outputResult, nil
}

func (documentModel DocumentModel) SearchDocumentByCountryIso(namespaceRequest *Entity.GlobalDocumentRequest) (data.TableGlobalDocument, error) {
	return data.TableGlobalDocument{}, nil
}
