package mail

import (
	"encoding/json"
	"test/utils"

	. "github.com/julvo/htmlgo"
	a "github.com/julvo/htmlgo/attributes"
)

const style string = `
    * {
      box-sizing: border-box;
      -webkit-box-sizing: border-box;
      -moz-box-sizing: border-box;
    }
    body {
      -webkit-font-smoothing: antialiased;
      background: white;
      padding: 2rem;
    }
    h2 {
      text-align: center;
      font-size: 18px;

      letter-spacing: 1px;
      color: black;
      padding: 30px 0;
    }
    a {
      text-align: center;
      font-size: 13px;
      margin: 1rem;
      font-weight: bold;
      text-decoration: none;
    }

    p {
      text-align: left;
      font-size: 13px;

      letter-spacing: 1px;
      color: black;
      font-weight: 500;
    }

    span {
      width: 100%;
      text-align: center;
    }

    .table-wrapper {
      margin: 10px 70px 70px 10px;
      box-shadow: 0px 35px 50px rgba(0, 0, 0, 0.2);
    }

    .fl-table {
      border-radius: 5px;
      font-size: 12px;
      font-weight: normal;
      border: none;
      border-collapse: collapse;
      width: 100%;
      max-width: 100%;
      white-space: nowrap;
      background-color: white;
    }

    .fl-table td,
    .fl-table th {
      text-align: center;
      padding: 8px;
    }

    .fl-table td {
      border-right: 1px solid #f8f8f8;
      font-size: 12px;
    }

    .fl-table td:first-child {
      border-left: 1px solid #f8f8f8;
    }

    .fl-table thead th {
      color: #ffffff;
      background: rgba(22, 160, 133, 1);
    }

    .fl-table thead th:nth-child(odd) {
      color: #ffffff;
      background: #324960;
    }

    .fl-table tr:nth-child(even) {
      background: #f8f8f8;
    }

    .fl-table tbody {
      border-bottom: #ffffff;
    }
`

func MakeTemplate(data interface{}) HTML {
  instance := make(map[string]interface{})
  t, _ := json.Marshal(data)
  err := json.Unmarshal(t, &instance)
  utils.CheckError(err)
  
	page :=
		Html5_(
			Head_(
				Meta(Attr(a.Charset_("utf-8"))),
				Meta(Attr(a.Name_("viewport"), a.Content_("width=device-width"), a.InitialScale_("1"))),
			),
			Body_(
				H2_(Text("확인하세요! WKMS가 DOWN 상태 입니다.")),
				P_("아래와 같이 서버 모니터링이 확인되었습니다. 본 메일은 회신이 불가능합니다."),
				P_("관리자 확인 필요합니다."),
				Div(Attr(a.Class_("table-wrapper")),
					Table(Attr(a.Class_("fl-table")),
						Thead_(
							Tr_(
								Th_("Timestamp"),
								Th_("Status"),
								Th_("Name"),
								Th_("Ip"),
								Th_("Port"),
								Th_("Zone"),
								Th_("Hostname")),
						),
						Tbody_(
							Tr_(
                Td_(Text(instance["timestamp"])),
                Td_(Text(instance["status"])),
                Td_(Text(instance["name"])),
                Td_(Text(instance["ip"])),
                Td_(Text(instance["port"])),
                Td_(Text(instance["zone"])),
                Td_(Text(instance["hostname"])),
	          	),
						),
					)),
				A(Attr(a.Href("https://wmp-siem.wemakeprice.work/app/uptime?dateRangeStart=now-24h&dateRangeEnd=now")), Text("보안팀 SIEM UPTIME 이동")),
				A(Attr(a.Href("http://10.102.181.45:9000/")), Text("보안팀 Jenkins 이동")),
			),
			Style_(Text(style)),
		)
	return page
}
