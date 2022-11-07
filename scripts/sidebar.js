const sidebar = document.getElementsByClassName("sidebar")[0];

const URL = "https://api.games.redbrick.dcu.ie";

function open() {
  sidebar.style.width = 250;
  sidebar.style.borderWidth = "0px 5px 0px 0px";
}

function close() {
  sidebar.style.width = 0;
  sidebar.style.borderWidth = 0;
}

function login() {
  let data = {
    "name": document.getElementById("username").value,
    "password": document.getElementById("passwd").value
  };

    let res = fetch(URL + "/login", {
      method: "POST",
      credentials: "include", body: JSON.stringify(data)
    }).then(async function (res) {
      if(!res.ok) {
        document.getElementById("badlogin").innerText = "Incorrect username or password!";
      }
      else {
        let me = await fetch(URL + "/me", {method: "GET", credentials: "include"});
        me = await me.json();

        let state;
        switch(me.state) {
          case 0:
            state = "Zombie (Core)";
            break;
          case 1:
            state = "Zombie";
            break;
          case 2:
            state = "Zombie (Core) (Stunned)";
            break;
          case 3:
            state = "Zombie (Stunned)";
            break;
          case 4:
            state = "Human (Infected)";
            break;
          case 5:
            state = "Human";
            break;
          case 6:
            state = "Dead";
            break;
        }

        document.getElementById("loglist").innerHTML = (
            "<li>Name: " + me.name + "</li>" +
            "<li>ID: " + me.id + "</li>" +
            "<li>State: " + state + "</li>" +
            "<li>Cures: " + me.cures + "</li>" +
            "<li>Revives: " + me.revives + "</li>"
        );

        document.getElementById("videoFlex").innerHTML = (
            "<span><p id='cur'>Cure</p><input id='cureId'><button type='button' id='cureBut'>Click!</button></span>" +
            "<span><p id='rev'>Revive</p><input id='revId'><button type='button' id='revBut'>Click!</button></span>" +
            "<span><p id='stu'>Stun</p><input id='stunId'><button type='button' id='stunBut'>Click!</button></span>" +
            "<span><p id='inf'>Infect</p><input id='infId'><button type='button' id='infBut'>Click!</button></span>"
        );

          document.getElementById("cureBut").addEventListener("click", cure);
          document.getElementById("revBut").addEventListener("click", rev);
          document.getElementById("stunBut").addEventListener("click", stun);
          document.getElementById("infBut").addEventListener("click", inf);
      }
    });
  }

  function cure() {
    let id = document.getElementById("cureId").value;
    fetch(URL + "/cure/" + id, {method: "GET", credentials: "include"})
        .then(function(res) {
          let elem = document.getElementById("cur");
          if(!res.ok) {
              switch(res.status) {
                  case 409:
                      elem.innerText = "Target is not a zombie";
                      break;
                  case 404:
                      elem.innerText = "Target does not exist";
                      break;
                  case 403:
                      elem.innerText = "No cures available";
                      break;
                  case 401:
                      elem.innerText = "User is not logged in";
                      break;
              }
          }
          else {
            elem.innerText = "Success!";
          }
        });
  }

function rev() {
  let id = document.getElementById("revId").value;
  fetch(URL + "/revive/" + id, {method: "GET", credentials: "include"})
      .then(function(res) {
        let elem = document.getElementById("rev");
        if(!res.ok) {
          switch(res.status) {
              case 409:
                  elem.innerText = "Target is not stunned";
                  break;
              case 404:
                  elem.innerText = "Target does not exist";
                  break;
              case 403:
                  elem.innerText = "No revives available";
                  break;
              case 401:
                  elem.innerText = "User is not logged in";
                  break;
          }
        }
        else {
          elem.innerText = "Success!";
        }
      });
}

function stun() {
    let id = document.getElementById("stunId").value;
    fetch(URL + "/kill/" + id, {method: "GET", credentials: "include"})
        .then(function(res) {
            let elem = document.getElementById("stu");
            if(!res.ok) {
                switch(res.status) {
                    case 409:
                        elem.innerText = "Target is not a zombie";
                        break;
                    case 404:
                        elem.innerText = "Target does not exist";
                        break;
                    case 403:
                        elem.innerText = "Killer is not a human";
                        break;
                    case 401:
                        elem.innerText = "User is not logged in";
                        break;
                }
            }
            else {
                elem.innerText = "Success!";
            }
        });
}

function inf() {
    let id = document.getElementById("infId").value;
    fetch(URL + "/tag/" + id, {method: "GET", credentials: "include"})
        .then(function(res) {
            let elem = document.getElementById("inf");
            if(!res.ok) {
                switch(res.status) {
                    case 409:
                        elem.innerText = "Target is not human";
                        break;
                    case 404:
                        elem.innerText = "Target does not exist";
                        break;
                    case 403:
                        elem.innerText = "Tagger is not a zombie";
                        break;
                    case 401:
                        elem.innerText = "User is not logged in";
                        break;
                }
            }
            else {
                elem.innerText = "Success!";
            }
        });
}

document.getElementById("openBut").addEventListener("click", open);
document.getElementById("closeBut").addEventListener("click", close);
document.getElementById("logBut").addEventListener("click", login);