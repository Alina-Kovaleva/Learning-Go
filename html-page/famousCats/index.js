const tiles = document.querySelectorAll(".game-tile");
let player = "x";
let winner = null;

const winningCombos = [
  [0, 1, 2],
  [3, 4, 5],
  [6, 7, 8],
  [0, 3, 6],
  [1, 4, 7],
  [2, 5, 8],
  [0, 4, 8],
  [2, 4, 6],
];

tiles.forEach((tile) => {
  tile.addEventListener("click", () => {
    if (
      !tile.classList.contains("tile-x") &&
      !tile.classList.contains("tile-o") &&
      winner === null
    ) {
      tile.classList.add(`tile-${player}`);
      if (checkForWinner()) {
        if (winner == "x") {
          showWinner(`Puss in Boots wins!`);
        } else if (winner == "o") {
          showWinner(`Matroskin wins!`);
        } else {
          showWinner("It's a tie!");
        }
      } else {
        player = player === "x" ? "o" : "x";
      }
    }
  });
});

function refreshButton() {
  tiles.forEach((tile) => {
    tile.classList.remove("tile-x", "tile-o");
  });
  player = "x";
  winner = null;
  showWinner("");
}

function checkForWinner() {
  for (let i = 0; i < winningCombos.length; i++) {
    const [a, b, c] = winningCombos[i];
    if (
      tiles[a].classList.contains(`tile-${player}`) &&
      tiles[b].classList.contains(`tile-${player}`) &&
      tiles[c].classList.contains(`tile-${player}`)
    ) {
      winner = player;
      return true;
    }
  }
  if (
    !winner &&
    Array.from(tiles).every(
      (tile) =>
        tile.classList.contains("tile-x") || tile.classList.contains("tile-o")
    )
  ) {
    winner = "tie";
    return true;
  }
  return false;
}

function showWinner(message) {
  const winnerMessage = document.querySelector(".winner-message");
  winnerMessage.textContent = message;
}

function toggleAudio(event) {
  if (event.target.classList.contains("garfield-quote-btn")) {
    var audio = document.getElementById("garfield-quote");
    var button = event.target;
    if (audio.paused) {
      audio.play();
      button.innerHTML = "Pause Garfield Quote";
    } else {
      audio.pause();
      button.innerHTML = "Play Garfield Quote";
    }
  } else {
    return false;
  }
}
