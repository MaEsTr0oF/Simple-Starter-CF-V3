// Подключение функционала "Чертогов Фрилансера"
import { isMobile } from "./functions.js";
// Подключение списка активных модулей
import { flsModules } from "./modules.js";

document.addEventListener('scroll', () => {
   const hiddenText = document.querySelectorAll('.off');
   hiddenText.forEach(element => {
      const rect = element.getBoundingClientRect();
      if (rect.top <= window.innerHeight && rect.bottom >= 0) {
         element.classList.add('on');
     }
   });
});

var slideIndex = 1;
showSlides(slideIndex);

function plusSlides(n) {
  showSlides(slideIndex += n);
}

function currentSlide(n) {
  showSlides(slideIndex = n);
}

function showSlides(n) {
  var i;
  var slides = document.getElementsByClassName('slide');
  var dots = document.getElementsByClassName('dot');

  if (n > slides.length) {
    slideIndex = 1;
  }
  if (n < 1) {
    slideIndex = slides.length;
  }
  for (i = 0; i < slides.length; i++) {
    slides[i].style.display = "none";

  }
  for (var i = 0; i < dots.length; i++) {
    dots[i].className = dots[i].className.replace("active", "");
  }
  slides[slideIndex - 1].style.display = "block";
  dots[slideIndex - 1].className += " active";
}
var i = 1;
  setInterval(function(){
  currentSlide(i)
  i++;
  if(i == 5)
  {
   i = 1;
  }
},4000);

document.getElementById('contactForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Отменяем стандартное поведение формы

  // Создаем объект FormData для получения данных из формы
  const formData = new FormData(this);

  // Получаем значения полей
  const nameSurename = formData.get('surename');
  const presenceChoice = formData.get('choosen');
  fetch("http:/localhost:8080/api", {
    method: "post",
    headers: {
      'Content-Type': 'application/json',
    },
    //make sure to serialize your JSON body
    body: JSON.stringify({
      FirstAndLastNames: nameSurename,
      attendance: presenceChoice,
    })
  })
  .then( (response) => {
     alert("ЖОПА");
  });  
});





