package model

import (
	"errors"
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
	"medusa-globalization-copywriting-system/cmd/handler/model/basic"
	"medusa-globalization-copywriting-system/tools/language"
	"time"
)

const (
	// ApplicationModelTableName tb_application
	documentTableName     = "tb_application_globalization_document_code"
	applicationId         = "application_id"
	namespaceId           = "namespace_id"
	documentCode          = "document_code"
	isEnable              = "is_enable"
	onlineTime            = "online_time"
	onlineOperatorUserId  = "online_operator_user_id"
	offlineTime           = "offline_time"
	offlineOperatorUserId = "offline_operator_user_id"
	offlineAccessUserId   = "offline_access_user_id"
	createTime            = "create_time"
	createUserId          = "create_user_id"
	deleteFlag            = "delete_flag"
	deleteTime            = "delete_time"
	deleteUserId          = "delete_user_id"

	documentValueTableName = "tb_application_globalization_document_value"
	documentId             = "document_id"
	countryIso             = "country_iso"
	countryName            = "country_name"
	documentValue          = "document_value"
	documentIsOnline       = "document_is_online"
	updateTime             = "update_time"
	updateUserId           = "update_user_id"
	lastUpdateDocument     = "last_update_document"
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
			applicationId: namespaceRequest.ApplicationId,
			namespaceId:   namespaceRequest.NamespaceId,
			onlineTime:    time.Now(),
			documentCode:  namespaceRequest.Key,
			createUserId:  0,
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
				documentId:    insertDocumentCodeResult,
				namespaceId:   namespaceRequest.NamespaceId,
				countryIso:    document.CountryCode,
				countryName:   languageCountry.CountryName,
				documentValue: document.LanguageValue,
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
