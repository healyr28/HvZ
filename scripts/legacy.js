const roman = ["", "", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV"]
const year = ["", "", 1988, 1990, 1991, 1992, 1994, 1997, 1999, 2000, 2001, 2002, 2006, 2009, 2013, 2016]
const console = ["", "", "Nintedo Entertainment System", "Nintedo Entertainment System", "Super Nintedo Entertainment System", "Super Nintedo Entertainment System", "Super Nintedo Entertainment System", "PlayStation", "PlayStation", "PlayStation", "PlayStation 2", "PlayStation 2 and PC", "PlayStation 2", "PlayStation 3", "PlayStation 3 and PC", "PlayStation 4"]

let slide = 2;
function changeSlides() {
  document.getElementById("videoFlex").innerHTML = "<p>Final Fantasy " + roman[slide] + "</p><img src=\"img/" + slide + ".png\"><p>Released in " + year[slide] + " for the " + console[slide] + "</p>";
  if(slide > 2) {
    const but = document.createElement("p");
	but.setAttribute("id", "left");
    but.innerHTML = "←";
    document.getElementById("videoFlex").appendChild(but);
    but.addEventListener("click", function() {
      slide -= 1;
      changeSlides();
    });
  }
  if(slide < 15) {
    const but = document.createElement("p");
	but.setAttribute("id", "right");
    but.innerHTML = "→";
    document.getElementById("videoFlex").appendChild(but);
    but.addEventListener("click", function() {
      slide += 1;
      changeSlides();
    });
  }
}

var player;
function onYouTubePlayerAPIReady() {
  player = new YT.Player("player", {
    events: {
      "onReady": onPlayerReady,
      "onStateChange": onPlayerStateChange
    }
  });
}

function onPlayerReady(event) {
  event.target.playVideo();
}

function onPlayerStateChange(event) {      
  if(event.data === 0) {
    changeSlides();
  }
}

document.getElementById("openSli").addEventListener("click", changeSlides);