// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	alarmThresholdManageFieldNames          = builder.RawFieldNames(&AlarmThresholdManage{})
	alarmThresholdManageRows                = strings.Join(alarmThresholdManageFieldNames, ",")
	alarmThresholdManageRowsExpectAutoSet   = strings.Join(stringx.Remove(alarmThresholdManageFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	alarmThresholdManageRowsWithPlaceHolder = strings.Join(stringx.Remove(alarmThresholdManageFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	alarmThresholdManageModel interface {
		Insert(ctx context.Context, data *AlarmThresholdManage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AlarmThresholdManageList, error)
		Update(ctx context.Context, data *AlarmThresholdManage) error
		Delete(ctx context.Context, id int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]AlarmThresholdManageList, error)
		DeleteSoft(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]AlarmThresholdManageList, error)
	}

	defaultAlarmThresholdManageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AlarmThresholdManage struct {
		Id           int64  `db:"id"`
		Name         string `db:"name"`           // 名称
		Config       string `db:"config"`         // 配置json
		Types        string `db:"types"`          //域名管理类型
		ProjectId    int64  `db:"project_id"`     // 项目id
		GameServerId int64  `db:"game_server_id"` // 游戏服id
		AssetId      int64  `db:"asset_id"`       // 主机id
		Level        int64  `db:"level"`          // 层级
		DelFlag      int64  `db:"del_flag"`       // 删除标记
		Remark       string `db:"remark"`         // 备注
	}

	AlarmThresholdManageList struct {
		Id              int64  `db:"id"`
		Name            string `db:"name"`              // 名称
		Config          string `db:"config"`            // 配置json
		Types           string `db:"types"`             //域名管理类型
		ProjectId       int64  `db:"project_id"`        // 项目id
		GameServerId    int64  `db:"game_server_id"`    // 游戏服id
		AssetId         int64  `db:"asset_id"`          // 主机id
		Level           int64  `db:"level"`             // 层级
		DelFlag         int64  `db:"del_flag"`          // 删除标记
		Remark          string `db:"remark"`            // 备注
		ProjectCn       string `db:"project_cn"`        // 备注
		Ips             string `db:"ips"`               // 备注
		GameServerAlias string `db:"game_server_alias"` // 备注

	}
)

var alarmMngSql = `SELECT
	%s 
FROM
	(
	SELECT
		alarm_threshold_manage.*,
		project.project_cn,
	IF
		( asset.outer_ip, CONCAT( '外:', asset.outer_ip, ' 内:', asset.inner_ip ), '' ) AS ips,
		IFNULL( game_server.server_alias, '' ) AS game_server_alias 
	FROM
		alarm_threshold_manage
		LEFT JOIN project ON alarm_threshold_manage.project_id = project.project_id
		LEFT JOIN asset ON alarm_threshold_manage.asset_id = asset.asset_id
		LEFT JOIN game_server ON alarm_threshold_manage.game_server_id = game_server.id 
		AND alarm_threshold_manage.project_id = game_server.project_id 
	) A 
WHERE
	del_flag = 0

%s
%s
`

func newAlarmThresholdManageModel(conn sqlx.SqlConn) *defaultAlarmThresholdManageModel {
	return &defaultAlarmThresholdManageModel{
		conn:  conn,
		table: "`alarm_threshold_manage`",
	}
}

func (m *defaultAlarmThresholdManageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAlarmThresholdManageModel) FindOne(ctx context.Context, id int64) (*AlarmThresholdManageList, error) {
	//query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", alarmThresholdManageRows, m.table)
	query := fmt.Sprintf(alarmMngSql, "*", "and id = ?", "limit 1")

	var resp AlarmThresholdManageList
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAlarmThresholdManageModel) Insert(ctx context.Context, data *AlarmThresholdManage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?,?, ?, ?, ?)", m.table, alarmThresholdManageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Config, data.Types, data.ProjectId, data.GameServerId, data.AssetId, data.Level, data.DelFlag, data.Remark)
	return ret, err
}

func (m *defaultAlarmThresholdManageModel) Update(ctx context.Context, data *AlarmThresholdManage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, alarmThresholdManageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Config, data.Types, data.ProjectId, data.GameServerId, data.AssetId, data.Level, data.DelFlag, data.Remark, data.Id)
	return err
}

func (m *defaultAlarmThresholdManageModel) tableName() string {
	return m.table
}

//分页条件查询
func (m *defaultAlarmThresholdManageModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]AlarmThresholdManageList, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(alarmMngSql, "*", condition, "limit ? offset ?")

	var resp []AlarmThresholdManageList
	err := m.conn.QueryRowsCtx(ctx, &resp, query, pageSize, offset)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//条件查询所有
func (m *defaultAlarmThresholdManageModel) FindAll(ctx context.Context, filters ...interface{}) (*[]AlarmThresholdManageList, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(alarmMngSql, "*", condition, "")

	var resp []AlarmThresholdManageList
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

//条件统计
func (m *defaultAlarmThresholdManageModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf(alarmMngSql, "count(*)", condition, "")

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

//软删除
func (m *defaultAlarmThresholdManageModel) DeleteSoft(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `del_flag`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, globalkey.DelStateYes, id)
	return err
}