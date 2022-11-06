const sidebar = document.getElementsByClassName("sidebar")[0];

function open() {
  sidebar.style.width = 250;
  sidebar.style.borderWidth = "0px 5px 0px 0px";
}

function close() {
  sidebar.style.width = 0;
  sidebar.style.borderWidth = 0;
}

document.getElementById("openBut").addEventListener("click", open);
document.getElementById("closeBut").addEventListener("click", close);
