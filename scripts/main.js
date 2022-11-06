const map = document.getElementsByClassName("map")[1]
const ship = document.getElementsByClassName("airship")[0]
let posLeft = 0
let posUp = 0
function randomEncounter() {
  console.log("You found a random encounter!!!!!!!!");
}

document.addEventListener("keydown", e => {
  switch(e.code) {
    case "ArrowRight":
    case "KeyD":
      if(posLeft > -4150) {posLeft -= 10;}
      ship.style.transform = "scale(1.5) rotateY(180deg)"
      map.style.left = posLeft + "px";
      break;
    case "ArrowLeft":
    case "KeyA":
      if(posLeft < 1020) {posLeft += 10;}
      ship.style.transform = "scale(1.5)";
      map.style.left =  posLeft + "px";
      break;
    case "ArrowUp":
    case "KeyW":
      if(posUp < 1020) {posUp += 10;}
      map.style.top = posUp + "px";
	  e.preventDefault();
      break;
    case "ArrowDown":
    case "KeyS":
      if(posUp > -4570) {posUp -= 10;}
      map.style.top = posUp + "px";
	  e.preventDefault();
      break;
    case "KeyR":
      randomEncounter();
      break;
  }
})
