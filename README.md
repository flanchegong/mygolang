# mygolang

```go
package main

import (
  "crypto/sha1"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/satori/go.uuid"
  "math/rand"
  "strconv"
  "strings"
  "time"
)

func main() {
  mysql()
}

//结构体
type Job struct {
  db    *sql.DB
  ch    chan int
  total int
  n     int
}

func mysql() {
  db, err := sql.Open("mysql", "root:root@/million?charset=utf8")
  if err != nil {
    fmt.Println("访问数据库出错", err)
    return
  }
  defer db.Close()
  db.SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
  db.SetMaxOpenConns(100)                  //设置最大连接数

  total := 5000
  gonum := 400
  fmt.Println("====start=====")
  start := time.Now()
  // 测试插入数据库的功能,每次最多同步20个工作协程
  jobChan := make(chan Job, 20)
  go worker(jobChan)
  //统计使用次数
  ch := make(chan int, gonum)
  for n := 0; n < gonum; n++ {
    job := Job{
      db:    db,
      ch:    ch,
      total: total,
      n:     n,
    }
    jobChan <- job
  }
  ii := 0
  for {
    <-ch
    ii++
    if ii >= gonum {
      break
    }
  }

  end := time.Now()
  curr := end.Sub(start)
  fmt.Println("run time:", curr)
}

func worker(jobChan <-chan Job) {
  for job := range jobChan {
    go insert(job)
  }
}

func insert(job Job) {
  var array [6]string = [6]string{"EL", "FMS", "GC", "OMS", "TEST", "UserCenter"}
  buf := make([]byte, 0, job.total)
  buf = append(buf, " insert into customer(customer_code,customer_firstname,customer_email,customer_currency,customer_telephone,customer_company_name,customer_reg_time,last_login_time,customer_update_time,company_code_sys,message_channel_key,financial_customer_code,customer_address,register_source,app_token,app_key) values "...)
  for i := 0; i < job.total; i++ {
    myid, _ := uuid.NewV4()
    customerCode := myid.String()
    customerFirstname := GetRandomString(8)
    customerEmail := "flanchegong@163.com"
    customerCurrency := "RMB"
    customerTelephone := "15811111111"
    customerCompanyName := "福建纵腾网络有限公司"
    //customerBillDate := time.Now().Format("2006-01-02 15:04:05")
    customerRegTime := "2006-01-02 15:04:05"
    lastLoginTime := "2019-10-08 15:04:05"
    customerUpdateTime := "2019-10-08 15:04:05"
    gongid, _ := uuid.NewV4()
    messageChannelKey := gongid.String()
    companyCodeSys := strconv.Itoa(rand.Intn(50000000))
    financialCustomerCode := "JR100"
    customerAddress := "深圳市民治向南4区"
    registerSource := array[rand.Intn(5)]
    //versions := time.Now().Format("2006-01-02 15:04:05")
    token, _ := uuid.NewV4()
    appToken := token.String()
    appKey :="flanche"// Sha1(appToken)

    //if i == job.total-1 {
    //  buf = append(buf, "'"+appKey+"');"...)
    //} else {
      buf = append(buf, "('"+customerCode+"',"...)
      buf = append(buf, "'"+customerFirstname+"',"...)
      buf = append(buf, "'"+customerEmail+"',"...)
      buf = append(buf, "'"+customerCurrency+"',"...)
      buf = append(buf, "'"+customerTelephone+"',"...)
      buf = append(buf, "'"+customerCompanyName+"',"...)
      // buf = append(buf, "('"+customerBillDate+"'),"...)
      buf = append(buf, "'"+customerRegTime+"',"...)
      buf = append(buf, "'"+lastLoginTime+"',"...)
      buf = append(buf, "'"+customerUpdateTime+"',"...)
      buf = append(buf, "'"+messageChannelKey+"',"...)
      buf = append(buf, "'"+companyCodeSys+"',"...)
      buf = append(buf, "'"+financialCustomerCode+"',"...)
      buf = append(buf, "'"+customerAddress+"',"...)
      buf = append(buf, "'"+registerSource+"',"...)
      // buf = append(buf, "('"+versions+"'),"...)
      buf = append(buf, "'"+appToken+"',"...)
      buf = append(buf, "'"+appKey+"'),"...)
   // }
  }
  ss := string(buf)
  ss = strings.TrimRight(ss, ",")
 // fmt.Println(ss)
//  runtime.Breakpoint()
  fmt.Println("第" + strconv.Itoa(job.n) + "次插入2.5万条数据！")
  _, err := job.db.Exec(ss)
  checkErr(err)
  fmt.Println("完成---" + strconv.Itoa(job.n) + "次插入2.5万条数据！")
  job.ch <- 1
}

func update(job Job) {
  var orderPlatformType [2]string = [2]string{"sale", "transfer"}
  var createType [3]string = [3]string{"hand", "excel","api"}
  var scCurrencyCode [6]string = [6]string{"AUD", "EUR", "GBP", "HKD", "RMB", "USD"}
  var platform [9]string = [9]string{"3DCART", "ALIEXPRESS", "AMAZON", "EBAY", "OTHER", "SHOPIFY","TONGTU","transfer","WISH"}
  var countryCode[42]string = [42]string{"AD","AF","AL","AT","AU","BA","BE","BR","BY","CA","CH","CN","CZ","DE","DK","DZ","ES","FR","GB","GU","IS","IT","JE","JP","KV","LI","MC","MD","ME","MK","MX","NL","NO","RS","RU","SI","SM","TR","TW","UA","US","VA"}
  var smCode[58]string = [58]string {"BGM","BOCI-FEDEXCIRI","BOCI-GROUD","BOCI-LARGER","BOCI-SMALL","CN-EMS","CZ_DHL_DOMESTICPK","CZ_DHL_INTLPK","DEDHL-PAKET","DHL-SMALLPARCEL","DHLSG","FA","FEDEX-LARGEPARCEL","FEDEX-PACKAGE","FEDEX-PACKAGE-B","FEDEX-SMALLPARCEL","FEDEX-SP","FEDEXG-RETURN","FEDEX_GROUPS","FEDEX_GROUPS_150","FEDEX_MC","FEDEX_OVERNIGHT","HERMES_DOMESTIC","HERMES_DOMESTIC_PC","HERMES_INTERNATIONAL","HERMES_LOCAL_BI","HY_DE_DHL","HY_EU_DHL","INT_E_PACKET","IPA_INT_ECONOMIC","JPEMS","JPTEST","LWEMS","RETURN-AUZ","RETURN-OR","ROYALMAIL-48H-G","ROYALMAIL-48H-G-S","TEST","TESTWL","TRACKED_48_NS","TRACKED_48_S","TUTU","UPS-GROUND","UPS-US-E","UPS-US-S","UPS_G","UPS_GROUND","USP-US-S","USPS-BPARCEL","USPS-LWPARCEL","USPS_LARGE_LETTER","WULIU001","WULIU002","XDP_TOW_MAN","XDP_UK","YESHANGQUAN","YODEL_48H","ZITI"}

  buf := make([]byte, 0, job.total)

  for i := 0; i < job.total; i++ {
    orderPlatformType := orderPlatformType[rand.Intn(2)]
    createType := createType[rand.Intn(3)]
    scCurrencyCode := scCurrencyCode[rand.Intn(6)]
    platform := platform[rand.Intn(9)]
    countryCode := countryCode[rand.Intn(42)]
    smCode := smCode[rand.Intn(58)]
    buf = append(buf, " UPDATE orders SET "...)
    buf = append(buf, " platform='"+platform+"',"...)
    buf = append(buf, " sc_currency_code='"+scCurrencyCode+"',"...)
    buf = append(buf, " order_platform_type='"+orderPlatformType+"',"...)
    buf = append(buf, " create_type='"+createType+"',"...)
    buf = append(buf, " country_code='"+countryCode+"',"...)
    buf = append(buf, " sm_code='"+smCode+"' WHERE order_id in (SELECT order_id FROM (SELECT order_id FROM orders ORDER BY order_id ASC  limit "+strconv.Itoa(i)+", 5000 ) AS tt); "...)
    ss := string(buf)
    fmt.Println("第" + strconv.Itoa(job.n) + "次修改5千条数据！")
    _, err := job.db.Exec(ss)
    checkErr(err)
    fmt.Println("完成---" + strconv.Itoa(job.n) + "次修改5千条数据！")
    job.ch <- 1
  }

}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func GetRandomString(l int) string {
  str := "0123456789abcdefghijklmnopqrstuvwxyz"
  bytes := []byte(str)
  result := []byte{}
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i := 0; i < l; i++ {
    result = append(result, bytes[r.Intn(len(bytes))])
  }
  return string(result)
}

func Sha1(s string) string {
  //产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
  h := sha1.New() // md5加密类似md5.New()
  //写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
  h.Write([]byte(s))
  //这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
  bs := h.Sum(nil)
  //SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
  abc, err := fmt.Printf("%x\n", bs)
  if err != nil {
    return string(abc)
  } else {
    return ""
  }

}
```
