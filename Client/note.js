var navBar = document.getElementById("navBar");
var newBtn = document.getElementById("newBtn");
var errReporter = document.getElementById("errReporter");

var title = document.getElementById("title");
var username = document.getElementById("username");
var writtenTime = document.getElementById("writtenTime");
var textCell = document.getElementById("textCell");

var searchTxt = document.getElementById("searchTxt");
var searchBtn = document.getElementById("searchBtn");

var saveBtn = document.getElementById("saveBtn");
var deleteBtn = document.getElementById("deleteBtn");
var content = document.getElementsByName("content");
var contentId = document.getElementById("contentId");
var textForm = document.getElementById("textForm");

const mainURL = "http://localhost:3002";

const errCase = {
  Server: "0",
  Form_Title: "1",
  Form_Username: "2",
};

const satisfyForm = {
  check: 0,
  do: 1,
};

//Cookie
function getCookie(name) {
  var value = document.cookie.match("(^|;) ?" + name + "=([^;]*)(;|$)");
  return value ? value[2] : null;
}

function deleteCookie(name) {
  document.cookie = name + "=; expires=Thu, 01 Jan 1999 00:00:10 GMT";
}

//Input
function errHandler(errCode) {
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
}

function satisfyFormInput(mode) {
  // mode check: only condition check, do: + temporary saving process
  if (title.value == "" || title.value.length > 255) {
    if (mode == satisfyForm.check) {
      errHandler(errCase.Form_Title);
      return false;
    } else {
      title.value = "temp";
    }
  } else if (username.value == "" || username.value.length > 10) {
    if (mode == satisfyForm.check) {
      errHandler(errCase.Form_Username);
      return false;
    } else {
      title.value = "Guest";
    }
  } else {
    return true;
  }
}

function isInContent(textCellBodyInHTML) {
  const isContent = textCellBodyInHTML == "" || textCellBodyInHTML == "<br>"; // true: 내용 없음
  return !isContent;
}

function initValues() {
  title.value = "";
  username.value = "";
  writtenTime.innerHTML = "";
  frames["textCell"].contentDocument.getElementsByTagName("body")[0].innerHTML =
    "";
  contentId.value = -1;
  content[0].value = "";
}

//Navigation
function createNavElement(navContents) {
  var index = 0;

  while (index < navContents.LEN) {
    var newLi = document.createElement("li");
    var newA = document.createElement("a");

    newLi.value = navContents.IDS[index];
    newLi.setAttribute("class", "navItem");

    newA.innerHTML = navContents.TITLES[index];
    newA.setAttribute("onclick", "clickNavItem(this)");
    newLi.appendChild(newA);
    navBar.appendChild(newLi);
    index++;
  }
}

function setNavContent() {
  const url = mainURL + "/get/nav/";

  fetch(url, { method: "GET" }).then(function (res) {
    res.json().then(function (navContents) {
      if (navContents != null) {
        createNavElement(navContents);
      }
    });
  });
}

// Update content
async function updateContent(textCellBody) {
  if (isInContent(textCellBody.innerHTML)) {
    // 내용 있음
    satisfyFormInput(satisfyForm.do);
    const saveURl = mainURL + "/save/";
    await fetch(saveURl, {
      method: "POST",
      cache: "no-cache",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body:
        "title=" +
        title.value +
        "&username=" +
        username.value +
        "&content=" +
        textCellBody.innerHTML +
        "&id=" +
        contentId.value,
    });
  }
}

// Get Content
async function clickNavItem(item) {
  const queryId = item.parentElement.value;
  var textCellBody =
    frames["textCell"].contentDocument.getElementsByTagName("body")[0];

  await updateContent(textCellBody);

  const getContentURL = mainURL + "/get/content?id=" + queryId;
  await fetch(getContentURL, { method: "GET" }).then(function (res) {
    res.json().then(function (sentContent) {
      title.value = sentContent.TITLE;
      username.value = sentContent.USERNAME;
      // writtenTime.innerHTML = sentContent.CREATED_TIME;
      // If you want to remove the digits after the decimal point...
      writtenTime.innerHTML = sentContent.CREATED_TIME.split(".")[0].replace(
        "T",
        " "
      );
      textCellBody.innerHTML = sentContent.CONTENT;
      contentId.value = queryId;
      content[0].value = "";
    });
  });
}

//Events
// 로드시 에러 쿠키가 구워져 있는지 확인함
window.addEventListener("load", function () {
  const errCode = getCookie("errorServer");
  errHandler(errCode);
  setNavContent();
  initValues();
  deleteCookie("errorServer");
});

// content 입력을 가능하게 하는 코드
textCell.addEventListener("mouseover", function () {
  textCell.contentDocument.designMode = "on";
});

// 입력 데이터 조건 검사 및 POST /save
saveBtn.addEventListener("click", function () {
  var textCellBody =
    frames["textCell"].contentDocument.getElementsByTagName("body")[0];

  if (satisfyFormInput(satisfyForm.check)) {
    content[0].value = textCellBody.innerHTML;
    textCellBody.innerHTML = "";

    textForm.action = "/save";
    textForm.submit();
  }
});

// 새로운 메모 만들기 == 그냥 저장하고 양식 비우기
newBtn.addEventListener("click", async function () {
  var textCellBody =
    frames["textCell"].contentDocument.getElementsByTagName("body")[0];

  await updateContent(textCellBody);

  initValues();
});

deleteBtn.addEventListener("click", function () {
  if (confirm("정말 삭제 하시겠습니까?")) {
    textForm.action = "/delete";
    textForm.submit();
  } else {
    return;
  }
});

searchBtn.addEventListener("click", function () {
  const queryUsername = searchTxt.value;
  var url = mainURL;

  if (queryUsername == "") {
    url += "/get/nav/";
  } else {
    url += "/get/nav?username=" + queryUsername;
  }

  fetch(url, { method: "GET" }).then((res) => {
    res.json().then((navContents) => {
      if (navContents != null) {
        var navItemAll = document.querySelectorAll(".navItem");
        navItemAll.forEach((navItemAllElem) => navItemAllElem.remove());
        
        if (navContents.LEN == 0) {
          var newLi = document.createElement("li");
          var newA = document.createElement("a");
          newLi.setAttribute("class", "navItem");
          newA.innerHTML = "검색 결과가 없습니다.";
          newLi.appendChild(newA);
          navBar.appendChild(newLi);
        }
        
        createNavElement(navContents)
      }
    });
  });
});
