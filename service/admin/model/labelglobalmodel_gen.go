// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/xfilters"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	labelGlobalFieldNames          = builder.RawFieldNames(&LabelGlobal{})
	labelGlobalRows                = strings.Join(labelGlobalFieldNames, ",")
	labelGlobalRowsExpectAutoSet   = strings.Join(stringx.Remove(labelGlobalFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	labelGlobalRowsWithPlaceHolder = strings.Join(stringx.Remove(labelGlobalFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	labelGlobalModel interface {
		Insert(ctx context.Context, data *LabelGlobal) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LabelGlobal, error)
		Delete(ctx context.Context, bindingId, labelId, projectId int64, resourceEn string) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]LabelGlobalAndLabel, error)
		FindAllBySearch(ctx context.Context, filters ...interface{}) (*[]SearchData, error)
		CountBySearch(ctx context.Context, filters ...interface{}) (int64, error)
		FindAllBindResource(ctx context.Context, filters ...interface{}) (*[]ViewSearchLabel, error)
		FindAllGroupByEnName(ctx context.Context, filters ...interface{}) (*[]LabelGlobalAndLabelExt, error)
		TransactInsert(ctx context.Context, in *adminclient.AddResourceReq) error
	}

	defaultLabelGlobalModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LabelGlobal struct {
		Id         int64  `db:"id"`
		LabelId    int64  `db:"label_id"`
		ResourceEn string `db:"resource_en"`
		BindingId  int64  `db:"binding_id"`
		ProjectId  int64  `db:"project_id"`
	}

	LabelGlobalAndLabel struct {
		Id         int64  `db:"id"`
		LabelId    int64  `db:"label_id"`
		ResourceEn string `db:"resource_en"`
		BindingId  int64  `db:"binding_id"`
		ProjectId  int64  `db:"project_id"`
		LabelType  string `db:"label_type"`
	}

	LabelGlobalAndLabelExt struct {
		Id         int64  `db:"id"`
		LabelId    int64  `db:"label_id"`
		ResourceEn string `db:"resource_en"`
		BindingId  int64  `db:"binding_id"`
		ProjectId  int64  `db:"project_id"`
		LabelType  string `db:"label_type"`
		BindingIds string `db:"binding_ids"`
	}

	SearchData struct {
		ViewDataContent string `db:"view_data_content"`
		ViewDataUrl     string `db:"view_data_url"`
		ViewJsonId      string `db:"view_json_id"`
		ViewTableName   string `db:"view_table_name"`
	}

	ViewSearchLabel struct {
		ViewResourceCnName  string `db:"view_resource_cn_name"`
		ViewResourceEnName  string `db:"view_resource_en_name"`
		ViewResourceRemark  string `db:"view_resource_remark"`
		ViewProjectId       int64  `db:"view_project_id"`
		ViewPrimaryKey      string `db:"view_primary_key"`
		ViewPrimaryKeyValue int64  `db:"view_primary_key_value"`
		ViewResourceType    string `db:"view_resource_type"`
		ViewResourceValue   string `db:"view_resource_value"`
		ViewDataContent     string `db:"view_data_content"`
		ViewDataUrl         string `db:"view_data_url"`
		ViewJsonId          string `db:"view_json_id"`
		ViewTableName       string `db:"view_table_name"`
		ViewShowCluster     int64  `db:"view_show_cluster"`
		ViewShowFeature     int64  `db:"view_show_feature"`
		ViewShowInstall     int64  `db:"view_show_install"`
		ViewShowOther       int64  `db:"view_show_other"`
		ViewSystemShow      int64  `db:"view_system_show"`
	}
)

func newLabelGlobalModel(conn sqlx.SqlConn) *defaultLabelGlobalModel {
	return &defaultLabelGlobalModel{
		conn:  conn,
		table: "`label_global`",
	}
}

func (m *defaultLabelGlobalModel) TransactInsert(ctx context.Context, in *adminclient.AddResourceReq) error {
	if len(in.ResourceData) != 0 {
		err := m.conn.Transact(func(session sqlx.Session) error {
			valueList := make([]string, 0)
			for _, data := range in.ResourceData {
				valueList = append(valueList, fmt.Sprintf("(%d,'%s',%d,%d)", data.LabelId, data.ResourceEn, data.BindingId, data.ProjectId))
			}

			insertsql := fmt.Sprintf("replace into %s (%s) values %s", m.table, labelGlobalRowsExpectAutoSet, strings.Join(valueList, " , "))
			stmt, err := session.Prepare(insertsql)
			if err != nil {
				return err
			}
			defer stmt.Close()

			// 返回任何错误都会回滚事务
			if _, err := stmt.ExecCtx(ctx); err != nil {
				return err
			}

			return nil
		})
		return err
	}
	return nil
}

func (m *defaultLabelGlobalModel) Insert(ctx context.Context, data *LabelGlobal) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?,?,?,?)", m.table, labelGlobalRowsExpectAutoSet)
	//fmt.Println(query)
	ret, err := m.conn.ExecCtx(ctx, query, data.LabelId, data.ResourceEn, data.BindingId, data.ProjectId)
	return ret, err
}

func (m *defaultLabelGlobalModel) FindOne(ctx context.Context, id int64) (*LabelGlobal, error) {
	query := fmt.Sprintf("select %s from %s where `label_id` = ? limit 1", labelGlobalRows, m.table)
	var resp LabelGlobal
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	//fmt.Println(err)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLabelGlobalModel) FindAll(ctx context.Context, filters ...interface{}) (*[]LabelGlobalAndLabel, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("SELECT l.label_type,lg.* FROM label l, label_global lg where l.label_id = lg.label_id %s", condition)
	var resp []LabelGlobalAndLabel
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLabelGlobalModel) FindAllGroupByEnName(ctx context.Context, filters ...interface{}) (*[]LabelGlobalAndLabelExt, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("SELECT l.label_type,lg.*,GROUP_CONCAT(binding_id) as binding_ids FROM label l, label_global lg where l.label_id = lg.label_id %s GROUP BY label_type,label_id,resource_en", condition)
	var resp []LabelGlobalAndLabelExt
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLabelGlobalModel) FindAllBindResource(ctx context.Context, filters ...interface{}) (*[]ViewSearchLabel, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " where " + xfilters.Xfilters(filters...)
	}

	query := fmt.Sprintf("select * from view_search_label %s", condition)
	//fmt.Println(query)
	var resp []ViewSearchLabel
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLabelGlobalModel) FindAllBySearch(ctx context.Context, filters ...interface{}) (*[]SearchData, error) {

	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " where " + xfilters.Xfilters(filters...)
	}

	query := fmt.Sprintf("select * from view_search_label %s ", condition)

	var resp []SearchData
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLabelGlobalModel) CountBySearch(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " where " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("select count(*) as count from view_search_all %s", condition)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultLabelGlobalModel) Delete(ctx context.Context, bindingValue, labelId, projectId int64, resourceEn string) error {
	var (
		query string
		err   error
	)
	if bindingValue == -1 {
		query = fmt.Sprintf("delete from %s where  `label_id` = ? and `resource_en` = ? and `project_id` = ?", m.table)
		_, err = m.conn.ExecCtx(ctx, query, labelId, resourceEn, projectId)
	} else {
		query = fmt.Sprintf("delete from %s where `binding_id` = ? and `label_id` = ? and `resource_en` = ? and `project_id` = ?", m.table)
		_, err = m.conn.ExecCtx(ctx, query, bindingValue, labelId, resourceEn, projectId)
	}

	return err
}

func (m *defaultLabelGlobalModel) tableName() string {
	return m.table
}