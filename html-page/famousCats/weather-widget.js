(function (d, s, id) {
  var js,
    fjs = d.getElementsByTagName(s)[0];
  if (!d.getElementById(id)) {
    js = d.createElement(s);
    js.id = id;
    js.src = "https://weatherwidget.io/js/widget.min.js";
    fjs.parentNode.insertBefore(js, fjs);
  }
})(document, "script", "weatherwidget-io-js");

const carousel = document.querySelector('.carousel-container');
const items = document.querySelectorAll('.carousel-item');
const prevBtn = document.querySelector('.carousel-prev');
const nextBtn = document.querySelector('.carousel-next');

let currentIndex = 0;

function showItem(index) {
  items.forEach((item, i) => {
    if (i === index) {
      item.style.display = 'block';
    } else {
      item.style.display = 'none';
    }
  });
}

function prevItem() {
  currentIndex--;
  if (currentIndex < 0) {
    currentIndex = items.length - 1;
  }
  showItem(currentIndex);
}

function nextItem() {
  currentIndex++;
  if (currentIndex >= items.length) {
    currentIndex = 0;
  }
  showItem(currentIndex);
}

prevBtn.addEventListener('click', prevItem);
nextBtn.addEventListener('click', nextItem);
showItem(currentIndex);