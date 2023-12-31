// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"

	"database/sql"
	"encoding/json"
	"fmt"

	"strings"
	"time"
	"ywadmin-v3/common/xtime"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	reportStreamMinuteFieldNames          = builder.RawFieldNames(&ReportStreamMinute{})
	reportStreamMinuteRows                = strings.Join(reportStreamMinuteFieldNames, ",")
	reportStreamMinuteRowsExpectAutoSet   = strings.Join(stringx.Remove(reportStreamMinuteFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	reportStreamMinuteRowsWithPlaceHolder = strings.Join(stringx.Remove(reportStreamMinuteFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	reportStreamMinuteModel interface {
		Insert(ctx context.Context, req *monitorclient.ReportAddReq) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ReportStreamMinute, error)
		Update(ctx context.Context, data *ReportStreamMinute) error
		Delete(ctx context.Context, id int64) error
		UpdateAssetHwInfo(ctx context.Context, req *monitorclient.ReportAddReq) error
		SelectListAllByType(ctx context.Context, params *monitorclient.GraphListReq) (map[string]interface{}, error)
		Count(ctx context.Context, assetId, reportTime int64, assetIp string) (int64, error)
		InsertReport(ctx context.Context, assetIp,remark string) (sql.Result, error)
	}

	defaultReportStreamMinuteModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ReportStreamMinute struct {
		Id          int64  `db:"id"`           // 主键ID
		AssetId     int64  `db:"asset_id"`     // 主机ID
		ReportType  string `db:"report_type"`  // 监控类型(disk,mem,net,cpu)
		ReportValue string `db:"report_value"` // 监控值
		Value       string `db:"value"`        // 值
		ReportTime  int64  `db:"report_time"`  // 上报时间
		State       int64  `db:"state"`
	}

	//监控数据cpu结构
	Cpu struct {
		CpuLoadOneMin     float32 `json:"cpu_load_one_min"`
		CpuLoadFiveMin    float32 `json:"cpu_load_five_min"`
		CpuLoadFifteenMin float32 `json:"cpu_load_fifteen_min"`
	}

	//监控数据mem结构
	Mem struct {
		FreeMem    float32 `json:"free_mem"`
		UsedMem    float32 `json:"used_mem"`
		TotalMem   float32 `json:"total_mem"`
		CacheMem   float32 `json:"cache_mem"`
		Percentage float32 `json:"percentage"`
	}

	//监控数据mem结构
	Disk struct {
		Name           string  `json:"name"`
		Util           float32 `json:"util"`
		Iowait         float32 `json:"iowait"`
		Percentage     float32 `json:"percentage"`
		FreeDiskSpace  float32 `json:"free_disk_space"`
		UsedDiskSpace  float32 `json:"used_disk_space"`
		TotalDiskSpace float32 `json:"total_disk_space"`
		MountPoint     string  `json:"mount_point"`
		DiskName       string  `json:"disk_name"`
	}

	//监控数据net结构
	Net struct {
		Name string  `json:"name"`
		Send float32 `json:"send"`
		Recv float32 `json:"recv"`
	}

	//解析监控数据格式
	CommonReportData struct {
		Cpu       Cpu    `json:"cpu"`
		Mem       Mem    `json:"mem"`
		Disk      []Disk `json:"disk"`
		Net       []Net  `json:"net"`
		AlarmTime int64  `json:"alarm_time"`
		Hostname  string `json:"hostname"`
		State     int64  `json:"state"`
	}
	//asset
	Asset struct {
		AssetId      int64  `db:"asset_id"`
		OuterIp      string `db:"outer_ip"`
		HardwareInfo string `db:"hardware_info"`
	}
	HardwareInfo struct {
		Hostname string `json:"hostname"`
	}

	//解析查询到cpu数据
	CpuData struct {
		Time  int64         `json:"time"`
		Tips  []interface{} `json:"tips"`
		Value float32       `json:"value"`
	}

	//解析查询到mem数据
	MemData struct {
		Time  int64         `json:"time"`
		Tips  []interface{} `json:"tips"`
		Value float32       `json:"value"`
	}

	//解析查询到disk数据
	DiskData struct {
		Name string   `json:"name"`
		Data DiskType `json:"data"`
	}

	DiskType struct {
		UsedDiskSpace []PublicData `json:"used_disk_space"`
		Util          []PublicData `json:"util"`
		Iowait        []PublicData `json:"iowait"`
	}

	PublicData struct {
		Time  int64         `json:"time"`
		Tips  []interface{} `json:"tips"`
		Value float32       `json:"value"`
	}

	//解析查询到net数据
	NetData struct {
		Name string  `json:"name"`
		Data NetType `json:"data"`
	}

	NetType struct {
		Send []PublicData `json:"send"`
		Recv []PublicData `json:"recv"`
	}
)

func newReportStreamMinuteModel(conn sqlx.SqlConn) *defaultReportStreamMinuteModel {
	return &defaultReportStreamMinuteModel{
		conn:  conn,
		table: "`report_stream_minute`",
	}
}

func (m *defaultReportStreamMinuteModel) InsertReport(ctx context.Context, assetIp,remark string) (sql.Result, error) {
	var (
		ret sql.Result
		err error
	)
	insertsql := []string{`INSERT INTO send_message.send_msg_record( msg_title, 
		msg_content, msg_type, msg_level, msg_to, access_ip, send_type,
		app_key, status, failure_log, create_date) VALUES ('%s', 
		'%s', 'feishu', 'disaster', '', '', '1', '10011', 1, '', %d);
	`, `INSERT INTO send_message.send_msg_record( msg_title, msg_content, msg_type, msg_level,
		 msg_to, access_ip, send_type, app_key, status, failure_log, create_date) VALUES ( '%s',
		 '%s' ,'wechat', 'disaster', 'jiayuanhao,zhangqingxian,dengxiguang', '', '1', '10011', 1, '', %d);
	`,
	}
	for i, v := range insertsql {
		title := fmt.Sprintf("[灾难]%s:%s最近5分钟丢失监控上报数据，请检查是否异常!", remark,assetIp)
		content := "报警时间：" + time.Now().Format("2006-01-02 15:04:05")
		newsql := fmt.Sprintf(v, title,
			content,
			time.Now().Unix(),
		)
		if i == 1 {
			newsql = fmt.Sprintf(v, title,
				title+"\n"+content,
				time.Now().Unix(),
			)
		}
		ret, err = m.conn.ExecCtx(ctx, newsql)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (m *defaultReportStreamMinuteModel) Count(ctx context.Context, assetId, reportTime int64, assetIp string) (int64, error) {
	query := "SELECT COUNT(*) count FROM report_stream_minute where asset_id=%d and report_time >= %d and report_time <= %d"
	query = fmt.Sprintf(query, assetId, reportTime-300, reportTime)
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

func (m *defaultReportStreamMinuteModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultReportStreamMinuteModel) FindOne(ctx context.Context, id int64) (*ReportStreamMinute, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", reportStreamMinuteRows, m.table)
	var resp ReportStreamMinute
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

func (m *defaultReportStreamMinuteModel) Insert(ctx context.Context, req *monitorclient.ReportAddReq) (sql.Result, error) {

	var (
		reportData CommonReportData
	)
	batch := make([]*ReportStreamMinute, 0)
	err := json.Unmarshal([]byte(req.Value), &reportData)
	if err != nil {
		return nil, err
	}
	var resp Asset
	ip := strings.Split(reportData.Hostname, ":")
	if len(ip) != 3 {
		return nil, errors.New("主机名错误，请检查")
	}
	err = m.conn.QueryRowCtx(ctx, &resp, fmt.Sprintf("select asset_id,outer_ip,hardware_info from asset where outer_ip = '%s'", ip[2]))
	if err != nil {
		return nil, err
	}

	for _, types := range []string{"disk", "net", "mem", "cpu"} {
		switch types {
		case "disk", "net":
			if types == "disk" {
				for _, v := range reportData.Disk {
					monitors, _ := json.Marshal(v)
					batch = append(batch, &ReportStreamMinute{
						AssetId:     resp.AssetId,
						ReportType:  types,
						Value:       v.Name,
						ReportValue: string(monitors),
						ReportTime:  reportData.AlarmTime,
						State:       reportData.State,
					})
				}
			} else {
				for _, v := range reportData.Net {
					monitors, _ := json.Marshal(v)
					batch = append(batch, &ReportStreamMinute{
						AssetId:     resp.AssetId,
						ReportType:  types,
						Value:       v.Name,
						ReportValue: string(monitors),
						ReportTime:  reportData.AlarmTime,
						State:       reportData.State,
					})
				}
			}
		case "mem", "cpu":
			var monitors []byte
			if types == "mem" {
				monitors, _ = json.Marshal(reportData.Mem)
			} else {
				monitors, _ = json.Marshal(reportData.Cpu)
			}
			batch = append(batch, &ReportStreamMinute{
				AssetId:     resp.AssetId,
				ReportType:  types,
				Value:       types,
				ReportValue: string(monitors),
				ReportTime:  reportData.AlarmTime,
				State:       reportData.State,
			})
		}
	}

	//开启批量事务
	err = m.conn.Transact(func(session sqlx.Session) error {
		tmp := make([]string, 0)
		for _, v := range batch {
			tmp = append(tmp, fmt.Sprintf("(%d,'%s','%s','%s',%d,%d)", v.AssetId, v.ReportType, v.ReportValue,
				v.Value, v.ReportTime, v.State))
		}
		insertsql := fmt.Sprintf("insert into %s (%s) values %s",
			m.table,
			reportStreamMinuteRowsExpectAutoSet,
			strings.Join(tmp, ","))
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
	if err != nil {
		return nil, err

	}
	return nil, err
}

func (m *defaultReportStreamMinuteModel) Update(ctx context.Context, data *ReportStreamMinute) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, reportStreamMinuteRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AssetId, data.ReportType, data.ReportValue, data.Value, data.ReportTime, data.Id)
	return err
}

func (m *defaultReportStreamMinuteModel) tableName() string {
	return m.table
}

func (m *defaultReportStreamMinuteModel) UpdateAssetHwInfo(ctx context.Context, req *monitorclient.ReportAddReq) error {
	var d HardwareInfo
	err := json.Unmarshal([]byte(req.Value), &d)
	if err != nil {
		return err
	}
	ip := strings.Split(d.Hostname, ":")
	if len(ip) != 3 {
		return errors.New("主机名错误，请检查")
	}
	query := fmt.Sprintf("update asset set `hardware_info` = '%s' where `outer_ip` = '%s'", req.Value, ip[2])
	_, err = m.conn.ExecCtx(ctx, query)
	return err
}

func (m *defaultReportStreamMinuteModel) SelectListAllByType(ctx context.Context, params *monitorclient.GraphListReq) (map[string]interface{}, error) {
	//当前时间分钟时间戳
	nowTime := xtime.GetMinTimeChangeSecondTime(gconv.String(time.Now().Unix()))

	cpudataonelist, cpudatafivelist, cpudatafifteenlist, memdatauselist, datadiskdataList, netdataList, err := m.SelectDataByLastOneHour(ctx, params, nowTime)
	if err != nil {
		return nil, err
	}
	pointNum, interval := 60, 60
	isHour := false
	if params.StartTime != 0 && params.EndTime != 0 && params.Granularity == "H" {
		pointNum = int(params.EndTime-params.StartTime) / 3600
		nowTime = params.EndTime
		interval = 3600
		isHour = true
	}
	if params.StartTime != 0 && params.EndTime != 0 && params.Granularity == "M" {
		pointNum = int(params.EndTime-params.StartTime) / 60
		nowTime = params.EndTime

	}
	diskTime := nowTime
	//cpu相关不足1小时数据，补齐
	if len(cpudataonelist) > 0 || len(cpudatafivelist) > 0 || len(cpudatafifteenlist) > 0 || len(memdatauselist) > 0 {
		var (
			//cpu
			cpudataoneinfo        []CpuData
			newcpudataoneinfo     []CpuData
			cpudatafiveinfo       []CpuData
			newcpudatafiveinfo    []CpuData
			cpudatafifteeninfo    []CpuData
			newcpudatafifteeninfo []CpuData
			//mem
			memdatauseinfo    []MemData
			newmemdatauseinfo []MemData

			//disk
			diskdatauseinfo    []DiskData
			newdiskdatauseinfo []DiskData

			//net
			netdatauseinfo    []NetData
			newnetdatauseinfo []NetData
		)
		marshalone, _ := json.Marshal(cpudataonelist)
		err1 := json.Unmarshal(marshalone, &cpudataoneinfo)
		marshalfive, _ := json.Marshal(cpudatafivelist)
		err2 := json.Unmarshal(marshalfive, &cpudatafiveinfo)
		marshalfifteen, _ := json.Marshal(cpudatafifteenlist)
		err3 := json.Unmarshal(marshalfifteen, &cpudatafifteeninfo)
		marshalmem, _ := json.Marshal(memdatauselist)
		err4 := json.Unmarshal(marshalmem, &memdatauseinfo)

		marshaldisk, _ := json.Marshal(datadiskdataList)
		err5 := json.Unmarshal(marshaldisk, &diskdatauseinfo)

		marshalnet, _ := json.Marshal(netdataList)
		err6 := json.Unmarshal(marshalnet, &netdatauseinfo)

		if err3 != nil || err1 != nil || err2 != nil || err4 != nil || err5 != nil || err6 != nil {
			fmt.Printf("datainfo parsing failed ,marshalone=%v,marshalfive=%v,marshalfifteen=%v,marshalmem=%v,marshaldisk=%v,marshalnet=%v\n", err1, err2, err3, err4, err5, err6)
		}
		//cpu
		cpudataoneinfo = (reverseSlice(cpudataoneinfo)).([]CpuData)
		cpudatafiveinfo = (reverseSlice(cpudatafiveinfo)).([]CpuData)
		cpudatafifteeninfo = (reverseSlice(cpudatafifteeninfo)).([]CpuData)
		//mem
		memdatauseinfo = (reverseSlice(memdatauseinfo)).([]MemData)
		// mem , cpu 初始化60个点
		nowTime = nowTime - 60
		for i := 1; i <= pointNum; i++ {
			newcpudataoneinfo = append(newcpudataoneinfo, CpuData{
				Time:  nowTime,
				Value: 0,
				Tips: []interface{}{
					"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
					0,
				},
			})
			newcpudatafiveinfo = append(newcpudatafiveinfo, CpuData{
				Time:  nowTime,
				Value: 0,
				Tips: []interface{}{
					"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
					0,
				},
			})
			newcpudatafifteeninfo = append(newcpudatafifteeninfo, CpuData{
				Time:  nowTime,
				Value: 0,
				Tips: []interface{}{
					"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
					0,
				},
			})
			newmemdatauseinfo = append(newmemdatauseinfo, MemData{
				Time:  nowTime,
				Value: 0,
				Tips: []interface{}{
					"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
					0,
				},
			})

			nowTime = nowTime - gconv.Int64(interval)

		}
		//disk 初始化60个点
		for _, vdisk := range diskdatauseinfo {
			usedDiskSpace := make([]PublicData, 0)
			util := make([]PublicData, 0)
			iowait := make([]PublicData, 0)
			nowTime = diskTime - 60
			for i := 1; i <= pointNum; i++ {
				usedDiskSpace = append(usedDiskSpace, PublicData{
					Time:  nowTime,
					Value: 0,
					Tips: []interface{}{
						"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
						0,
					},
				})
				util = append(util, PublicData{
					Time:  nowTime,
					Value: 0,
					Tips: []interface{}{
						"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
						0,
					},
				})
				iowait = append(iowait, PublicData{
					Time:  nowTime,
					Value: 0,
					Tips: []interface{}{
						"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
						0,
					},
				})
				nowTime = nowTime - gconv.Int64(interval)
			}

			newdiskdatauseinfo = append(newdiskdatauseinfo, DiskData{
				Name: vdisk.Name,
				Data: DiskType{
					UsedDiskSpace: usedDiskSpace,
					Util:          util,
					Iowait:        iowait,
				},
			})

		}

		//net 初始化60个点
		for _, vdisk := range netdatauseinfo {
			send := make([]PublicData, 0)
			recv := make([]PublicData, 0)
			nowTime = diskTime - 60
			for i := 1; i <= pointNum; i++ {
				send = append(send, PublicData{
					Time:  nowTime,
					Value: 0,
					Tips: []interface{}{
						"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
						0,
					},
				})
				recv = append(recv, PublicData{
					Time:  nowTime,
					Value: 0,
					Tips: []interface{}{
						"<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(nowTime),
						0,
					},
				})

				nowTime = nowTime - gconv.Int64(interval)
			}
			newnetdatauseinfo = append(newnetdatauseinfo, NetData{
				Name: vdisk.Name,
				Data: NetType{
					Send: send,
					Recv: recv,
				},
			})

		}

		//补齐近一小时数据
		//cpu one
		for i2, v2 := range newcpudataoneinfo {
			for _, v1 := range cpudataoneinfo {
				if isHour {
					fmt.Println(v2.Time)
					fmt.Println(v1.Time)
					fmt.Println(v2.Time - v1.Time)
					fmt.Println("++++++++++++++++")
					if v2.Time-v1.Time < 3600 && v2.Time-v1.Time >= 0 {
						newcpudataoneinfo[i2] = v1
					}
				} else {
					if v2.Time-v1.Time > -60 && v2.Time-v1.Time <= 0 {
						newcpudataoneinfo[i2] = v1
					}
				}

			}
		}
		//cpu five
		for i2, v2 := range newcpudatafiveinfo {
			for _, v1 := range cpudatafiveinfo {

				if isHour {
					if v2.Time-v1.Time < 3600 && v2.Time-v1.Time >= 0 {
						newcpudatafiveinfo[i2] = v1
					}
				} else {
					if v2.Time-v1.Time > -60 && v2.Time-v1.Time <= 0 {
						newcpudatafiveinfo[i2] = v1
					}
				}
			}
		}
		//cpu fifteen
		for i2, v2 := range newcpudatafifteeninfo {
			for _, v1 := range cpudatafifteeninfo {

				if isHour {
					if v2.Time-v1.Time < 3600 && v2.Time-v1.Time >= 0 {
						newcpudatafifteeninfo[i2] = v1
					}
				} else {
					if v2.Time-v1.Time > -60 && v2.Time-v1.Time <= 0 {
						newcpudatafifteeninfo[i2] = v1
					}
				}
			}
		}
		//mem
		for i2, v2 := range newmemdatauseinfo {
			for _, v1 := range memdatauseinfo {

				if isHour {
					if v2.Time-v1.Time < 3600 && v2.Time-v1.Time >= 0 {
						newmemdatauseinfo[i2] = v1
					}
				} else {
					if v2.Time-v1.Time > -60 && v2.Time-v1.Time <= 0 {
						newmemdatauseinfo[i2] = v1
					}
				}
			}
		}
		//disk
		for _, v1 := range newdiskdatauseinfo {
			for _, v2 := range diskdatauseinfo {
				for i111, v111 := range v1.Data.UsedDiskSpace {
					for _, v222 := range v2.Data.UsedDiskSpace {

						if isHour {
							if v111.Time-v222.Time < 3600 && v111.Time-v222.Time >= 0 && v1.Name == v2.Name {
								v1.Data.UsedDiskSpace[i111] = v222
							}

						} else {
							if v111.Time-v222.Time > -60 && v111.Time-v222.Time <= 0 && v1.Name == v2.Name {
								v1.Data.UsedDiskSpace[i111] = v222
							}
						}

					}
				}

				for i111, v111 := range v1.Data.Util {
					for _, v222 := range v2.Data.Util {

						if isHour {
							if v111.Time-v222.Time < 3600 && v111.Time-v222.Time >= 0 && v1.Name == v2.Name {
								v1.Data.Util[i111] = v222
							}

						} else {
							if v111.Time-v222.Time > -60 && v111.Time-v222.Time <= 0 && v1.Name == v2.Name {
								v1.Data.Util[i111] = v222
							}
						}
					}
				}

				for i111, v111 := range v1.Data.Iowait {
					for _, v222 := range v2.Data.Iowait {

						if isHour {
							if v111.Time-v222.Time < 3600 && v111.Time-v222.Time >= 0 && v1.Name == v2.Name {
								v1.Data.Iowait[i111] = v222
							}

						} else {
							if v111.Time-v222.Time > -60 && v111.Time-v222.Time <= 0 && v1.Name == v2.Name {
								v1.Data.Iowait[i111] = v222
							}
						}
					}
				}

			}
			v1.Data.UsedDiskSpace = reverseSlice(v1.Data.UsedDiskSpace).([]PublicData)
			v1.Data.Util = reverseSlice(v1.Data.Util).([]PublicData)
			v1.Data.Iowait = reverseSlice(v1.Data.Iowait).([]PublicData)

		}
		//net
		for _, v1 := range newnetdatauseinfo {
			for _, v2 := range netdatauseinfo {
				for _, v22 := range v2.Data.Send {
					for i12, v12 := range v1.Data.Send {

						if isHour {
							if v12.Time-v22.Time < 3600 && v12.Time-v22.Time >= 0 && v2.Name == v1.Name {
								v1.Data.Send[i12] = v22
							}

						} else {
							if v12.Time-v22.Time > -60 && v12.Time-v22.Time <= 0 && v2.Name == v1.Name {
								v1.Data.Send[i12] = v22
							}
						}

					}
				}
				for _, v23 := range v2.Data.Recv {
					for i13, v13 := range v1.Data.Recv {

						if isHour {
							if v13.Time-v23.Time < 3600 && v13.Time-v23.Time >= 0 && v2.Name == v1.Name {
								v1.Data.Recv[i13] = v23
							}

						} else {
							if v13.Time-v23.Time > -60 && v13.Time-v23.Time <= 0 && v2.Name == v1.Name {
								v1.Data.Recv[i13] = v23
							}
						}
					}

				}

			}

			v1.Data.Send = reverseSlice(v1.Data.Send).([]PublicData)
			v1.Data.Recv = reverseSlice(v1.Data.Recv).([]PublicData)

		}

		//最后将cpu数据重新赋值
		cpudataonelist = reverseSlice(newcpudataoneinfo).([]CpuData)
		cpudatafivelist = reverseSlice(newcpudatafiveinfo).([]CpuData)
		cpudatafifteenlist = reverseSlice(newcpudatafifteeninfo).([]CpuData)
		//最后将mem数据重新赋值
		memdatauselist = reverseSlice(newmemdatauseinfo).([]MemData)
		//最后将disk数据重新赋值
		datadiskdataList = newdiskdatauseinfo
		//最后将net数据重新赋值
		netdataList = newnetdatauseinfo
	} else {
		cpudataonelist = make([]CpuData, 0)
		cpudatafivelist = make([]CpuData, 0)
		cpudatafifteenlist = make([]CpuData, 0)
		memdatauselist = make([]MemData, 0)
		datadiskdataList = make([]DiskData, 0)
		netdataList = make([]NetData, 0)
	}

	//结果字典
	resultMap := make(map[string]interface{}, 0)
	//cpu
	resultMap["cpu"] = []interface{}{
		map[string]interface{}{
			"name": "1分钟负载",
			"data": cpudataonelist,
		},
		map[string]interface{}{
			"name": "5分钟负载",
			"data": cpudatafivelist,
		},
		map[string]interface{}{
			"name": "15分钟负载",
			"data": cpudatafifteenlist,
		},
	}
	//map
	resultMap["mem"] = []interface{}{
		map[string]interface{}{
			"name": "已用内存",
			"data": memdatauselist,
		},
	}
	resultMap["disk"] = datadiskdataList
	resultMap["net"] = netdataList

	return resultMap, nil
}

//切片倒序，按照指定的
func reverseSlice(source interface{}) (desc interface{}) {
	switch source.(type) {
	case []CpuData:
		s := source.([]CpuData)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	case []MemData:
		s := source.([]MemData)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	case []PublicData:
		s := source.([]PublicData)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s

	}

	return nil
}

func (m *defaultReportStreamMinuteModel) SelectListAll(query string) (*[]ReportStreamMinute, error) {
	var resp []ReportStreamMinute
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultReportStreamMinuteModel) SelectDataByLastOneHour(ctx context.Context, params *monitorclient.GraphListReq, nowTime int64) (cpudataonelist, cpudatafivelist, cpudatafifteenlist []CpuData,
	memdatauselist []MemData, datadiskdataList []DiskData, netdataList []NetData, _ error) {
	var (
		allgroupSql string
		allSql      string
	)
	if params.Granularity == "H" && params.EndTime != 0 && params.StartTime != 0 {
		//当前组数据
		allgroupSql = fmt.Sprintf("SELECT * FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) GROUP BY value ", params.AssetId, params.StartTime, params.EndTime)
		//当前所有数据
		allSql = fmt.Sprintf("SELECT id,asset_id,report_type,report_value,value,hours,ROUND(unix_timestamp(hours)) report_time,state"+
			" FROM (SELECT *,FROM_UNIXTIME(report_time,'%%Y-%%m-%%d-%%H:00:00') hours FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) ORDER BY report_time) A GROUP BY value,hours", params.AssetId, params.StartTime, params.EndTime)
	} else if params.Granularity == "M" && params.EndTime != 0 && params.StartTime != 0 {
		//近4小时
		allgroupSql = fmt.Sprintf("SELECT * FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) GROUP BY value ", params.AssetId, params.StartTime, params.EndTime)
		//近4小时所有数据
		allSql = fmt.Sprintf("SELECT * FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) ORDER BY report_time", params.AssetId, params.StartTime, params.EndTime)
	} else {
		//一小时之前
		allgroupSql = fmt.Sprintf("SELECT * FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) GROUP BY value ", params.AssetId, nowTime-3600, nowTime)
		//一小时之前所有数据
		allSql = fmt.Sprintf("SELECT * FROM report_stream_minute WHERE asset_id=%d "+
			"and ( report_time >= %d and report_time <= %d ) ORDER BY report_time", params.AssetId, nowTime-3600, nowTime)
	}

	allgroups, err := m.SelectListAll(allgroupSql)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	allData, err2 := m.SelectListAll(allSql)
	if err2 != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	//处理数据--
	for _, vgroup := range *allgroups {
		//disk分组
		diskuseList := make([]PublicData, 0)
		diskutilList := make([]PublicData, 0)
		diskwaitList := make([]PublicData, 0)
		//net分组
		netsendList := make([]PublicData, 0)
		netrecvList := make([]PublicData, 0)
		for _, vall := range *allData {
			if vgroup.ReportType == vall.ReportType {
				switch vall.ReportType {
				case "disk", "net":
					if vall.Value == vgroup.Value && vall.ReportType == "disk" {
						var disk Disk
						json.Unmarshal([]byte(vall.ReportValue), &disk)
						//usePercent := strings.Split(disk.Percentage, "%")
						//useSpace, _ := strconv.ParseFloat(usePercent[0], 32)
						tmpUtil := disk.Util
						tmpIowait := disk.Iowait
						useSpace := disk.Percentage
						tmperror := xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
						if vall.State == 2 {
							useSpace = -1
							tmpUtil = -1
							tmpIowait = -1
							tmperror = "<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))

						}
						diskuseList = append(diskuseList, PublicData{
							Tips: []interface{}{
								tmperror,
								fmt.Sprintf("总磁盘：%.2fGB", disk.TotalDiskSpace),
								fmt.Sprintf("空闲磁盘：%.2fGB", disk.FreeDiskSpace),
								fmt.Sprintf("已使用磁盘：%.2fGB", disk.UsedDiskSpace),
								fmt.Sprintf("磁盘使用率：%.2f%%", disk.Percentage),
							},
							Value: useSpace,
							Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
						diskutilList = append(diskutilList, PublicData{
							Tips: []interface{}{
								tmperror,
								disk.Util,
							},
							Value: tmpUtil,
							Time:  vall.ReportTime})
						diskwaitList = append(diskwaitList, PublicData{
							Tips: []interface{}{
								tmperror,
								disk.Iowait,
							},
							Value: tmpIowait,
							Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
					}
					if vall.Value == vgroup.Value && vall.ReportType == "net" {

						var net Net
						json.Unmarshal([]byte(vall.ReportValue), &net)
						tmpSend := net.Send
						tmpRecv := net.Recv
						tmperror := xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
						if vall.State == 2 {
							tmpSend = -1
							tmpRecv = -1
							tmperror = "<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
						}
						netsendList = append(netsendList, PublicData{
							Tips: []interface{}{
								tmperror,
								net.Send,
							},
							Value: tmpSend,
							Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
						netrecvList = append(netrecvList, PublicData{
							Tips: []interface{}{
								tmperror,
								net.Recv,
							},
							Value: tmpRecv,
							Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})

					}

				case "mem":
					var mem Mem
					json.Unmarshal([]byte(vall.ReportValue), &mem)
					usermem := mem.UsedMem
					tmperror := xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
					if vall.State == 2 {
						usermem = -1
						tmperror = "<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
					}
					memdatauselist = append(memdatauselist, MemData{

						Tips: []interface{}{
							tmperror,
							fmt.Sprintf("总内存：%.2fMB", mem.TotalMem),
							fmt.Sprintf("空内存：%.2fMB", mem.FreeMem),
							fmt.Sprintf("已使用内存：%.2fMB", mem.UsedMem),
						},
						Value: float32(usermem),
						Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})

				case "cpu":
					var cpu Cpu
					json.Unmarshal([]byte(vall.ReportValue), &cpu)
					tmpCpuLoadOneMin := cpu.CpuLoadOneMin
					tmpCpuLoadFiveMin := cpu.CpuLoadFiveMin
					tmpCpuLoadFifteenMin := cpu.CpuLoadFifteenMin
					tmperror := xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
					if vall.State == 2 {
						tmperror = "<span style=\"color:red\">[异常]</span>" + xtime.GetTimetampByTimeMinu(xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime)))
						tmpCpuLoadOneMin, tmpCpuLoadFiveMin, tmpCpuLoadFifteenMin = -1, -1, -1
					}
					cpudataonelist = append(cpudataonelist, CpuData{
						Tips:  []interface{}{tmperror, cpu.CpuLoadOneMin},
						Value: tmpCpuLoadOneMin,
						Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
					cpudatafivelist = append(cpudatafivelist, CpuData{
						Tips:  []interface{}{tmperror, cpu.CpuLoadFiveMin},
						Value: tmpCpuLoadFiveMin,
						Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
					cpudatafifteenlist = append(cpudatafifteenlist, CpuData{
						Tips:  []interface{}{tmperror, cpu.CpuLoadFifteenMin},
						Value: tmpCpuLoadFifteenMin,
						Time:  xtime.GetMinTimeChangeSecondTime(gconv.String(vall.ReportTime))})
				}

			}
		}

		if vgroup.ReportType == "disk" && len(diskuseList) > 0 && len(diskutilList) > 0 && len(diskwaitList) > 0 {
			datadiskdataList = append(datadiskdataList, DiskData{
				Name: vgroup.Value,
				Data: DiskType{
					UsedDiskSpace: diskuseList,
					Util:          diskutilList,
					Iowait:        diskwaitList,
				},
			})
		}

		if vgroup.ReportType == "net" && len(netsendList) > 0 && len(netrecvList) > 0 {

			netdataList = append(netdataList, NetData{
				Name: vgroup.Value,
				Data: NetType{
					Send: netsendList,
					Recv: netrecvList,
				},
			})
		}

	}

	return
}
