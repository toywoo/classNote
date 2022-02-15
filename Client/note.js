var errReporter = document.getElementById("errReporter");

var title = document.getElementById("title");
var username = document.getElementById("username");
var textCell = document.getElementById("textCell");

var saveBtn = document.getElementById("saveBtn");
var content = document.getElementsByName("content");
var textForm = document.getElementById("textForm");

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
