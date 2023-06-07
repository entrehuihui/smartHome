import axios from "axios";
import router from "@/router";
import * as XLSX from "xlsx";
// --------------- //
var logininfo = "";

function getAuthorization() {
  if (!logininfo) {
    getCookie();
  }
  return {
    headers: {
      authorization:logininfo,
      // 'Cache-Control': 'no-cache',
    },
    // httpsAgent: new https.Agent({
    //   rejectUnauthorized: false,
    // }),
  };
}

function ret(response) {
  if (response == undefined) {
    response = {};
    response.data = {};
    response.data.code = 500;
    response.data.error = "连接服务器失败,请稍后重试";
    response.data.message = "连接服务器失败,请稍后重试";
  } else if (response.status == 200) {
    // --------
  } else if (response.status == 301) {
    response.data.error = "登录超时";
    response.data.message = "登录超时";
    response.data.code = 301;
    router.push("/");
  } else if (response.status == 500) {
    var msg = response.data;
    if (msg.code == 301) {
      router.push("/");
      return response.data;
    }
  } else {
    response.data = {};
    response.data.error = "网络错误";
    response.data.message = "网络错误";
    response.data.code = 999;
  }
  if (response.data.code == 301) {
    router.push("/");
  }
  return response.data;
}
async function get(url, data) {
  if (data) {
    const datalist = Object.entries(data);
    let query = "?";
    for (const da of datalist) {
      query += da[0] + "=" + da[1] + "&";
    }
    url += query;
  }
  
  var response = await axios.get(url, getAuthorization()).catch((error) => {
    return error.response;
  });
  return ret(response);
}

async function post(url, data) {
  var response = await axios
    .post(url, data, getAuthorization())
    .catch((error) => {
      return error.response;
    });
  return ret(response);
}

async function put(url, data) {
  var response = await axios
    .put(url, data, getAuthorization())
    .catch((error) => {
      return error.response;
    });
  return ret(response);
}

async function del(url,paprm) {
  url = url +'/' + paprm
  var response = await axios.delete(url, getAuthorization()).catch((error) => {
    return error.response;
  });
  return ret(response);
}

function checkEmail(value) {
  var ePattern = /^([A-Za-z0-9_\-.])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})$/;
  if (ePattern.test(value)) {
    return true;
  }
  return false;
}

// 校验登录名：只能输入4-20个以字母开头、可带数字、“_”、“.”的字串
function checkUser(value) {
  var uPattern = /^[a-zA-Z]{1}([a-zA-Z0-9]|[._-]){3,19}$/;
  //输出 true
  return uPattern.test(value);
}

// 密码强度正则，最少6位，包括至少1个大写字母，1个小写字母，1个数字
function checPassword(value) {
  var pPattern = /^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).*$/;
  //输出 true
  return pPattern.test(value);
}

function checkName(value) {
  // ^[\u4E00-\u9FA5A-Za-z0-9_]+$
  var pPattern = /^[@.\u4E00-\u9FA5A-Za-z0-9_-]+$/;
  //输出 true
  return pPattern.test(value);
}

// 获取登陆cookie
function getCookie() {
  var key = cookieName;
  if (document.cookie.length == 0) {
    return "";
  }
  var start = document.cookie.indexOf(key + "=");
  if (start == -1) {
    return "";
  }
  start = start + key.length + 1;
  var end = document.cookie.indexOf(";", start);
  if (end === -1) end = document.cookie.length;
  let jwts = unescape(document.cookie.substring(start, end));
  if (jwts.length < 30) {
    setCookie("", -1);
    return "";
  }
  var jwtJson = JSON.parse(jwts)
  logininfo = jwtJson.token;
  return jwtJson;
}

const cookieName = "smnarkhome";

function setCookie(userinfo) {
  if (!userinfo ) {
    document.cookie = cookieName + "=" + ";SameSite=Lax";
    return;
  }
  logininfo = userinfo.token;
  var cookie = JSON.stringify(userinfo)
  document.cookie = cookieName + "=" + escape(cookie) + ";SameSite=Lax";
}

async function getTime() {
  return parseInt(new Date().getTime()) + '';
}

const exportExcel = (name, selectedRows, optionsValue) => {
  if (selectedRows.length == 0) {
    return;
  }
  // data.selectedRows.map((v) => {
  //     let obj = {};
  //     obj.商户编号 = v.ID;
  //     obj.名称 = v.name;
  //     obj.支付宝账户 = v.alipay;
  //     obj.邮箱 = v.emails;
  //     obj.地址 = v.addr;
  //     obj.创建时间 = new Date(parseInt(v.createtime)).toLocaleString();
  //     obj.备注 = v.details;
  //     arr.push(obj);
  // });
  const ws = XLSX.utils.json_to_sheet(selectedRows);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "商户信息");

  var options = {
    "!cols": Object.keys(selectedRows[0]).map(() => {
      return { wpx: 150 };
    }),
  };
  if (optionsValue) {
    options = {
      "!cols": optionsValue,
    };
  }
  ws["!cols"] = options["!cols"];
  // XLSX.writeFile(wb, "商户信息.xlsx");
  name += ".xlsx";
  XLSX.writeFile(wb, name);
};

export default {
  logininfo,
  post,
  get,
  put,
  del,
  checkUser,
  checPassword,
  getTime,
  checkEmail,
  setCookie,
  getAuthorization,
  getCookie,
  checkName,
  exportExcel,
};
