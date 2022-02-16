var errReporter = document.getElementById("errReporter");

var title = document.getElementById("title");
var username = document.getElementById("username");
var textCell = document.getElementById("textCell");

var saveBtn = document.getElementById("saveBtn");
var content = document.getElementsByName("content");
var textForm = document.getElementById("textForm");

const errCase = {
  Server: "0",
  Form_Title: "1",
  Form_Username: "2",
};

function getCookie(name) {
  var value = document.cookie.match("(^|;) ?" + name + "=([^;]*)(;|$)");
  return value ? value[2] : null;
}

function deleteCookie(name) {
  document.cookie = name + "=; expires=Thu, 01 Jan 1999 00:00:10 GMT";
}

// 로드시 에러 쿠키가 구워져 있는지 확인함
window.addEventListener("load", function () {
  const errCode = getCookie("errorServer");
  switch (errCode) {
    case errCase.Server:
      errReporter.innerText = "서버에 오류가 발생했습니다.";
      break;

    case errCase.Form_Title:
      errReporter.innerText =
        "제목이 255자 보다 길거나 없습니다. 다시 입력 해주세요.";
      break;

    case errCase.Form_Username:
      errReporter.innerText =
        "글쓴이의 글자 길이가 10자 보다 길거나 없습니다. 다시 입력 해주세요.";
      break;

    default:
      errReporter.innerText = "";
      break;
  }
  deleteCookie("errorServer");
});

// content 입력을 가능하게 하는 코드
textCell.addEventListener("mouseover", function () {
  textCell.contentDocument.designMode = "on";
});

// 입력 데이터 조건 검사 및 POST /save
saveBtn.addEventListener("click", function () {
  var textCellBody = textCell.contentDocument.getElementsByTagName("body")[0];

  if (title.value == "" || title.value.length > 255) {
    errReporter.innerText =
      "제목이 255자 보다 길거나 없습니다. 다시 입력 해주세요.";
  } else if (username.value == "" || username.value.length > 10) {
    errReporter.innerText =
      "글쓴이의 글자 길이가 10자 보다 길거나 없습니다. 다시 입력 해주세요.";
  } else {
    content[0].value = textCellBody.innerHTML;
    textCellBody.innerHTML = "";

    textForm.action = "/save";
    textForm.submit();
  }
});
