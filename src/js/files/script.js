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


document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('form');

  form.addEventListener('submit', async function(e) {
      e.preventDefault();

      let formData = new FormData(form);

      try {
          let response = await fetch('sendmail.php', {
              method: 'POST',
              body: formData
          });

          if (response.ok) {
              const result = await response.json(); // Получаем JSON-ответ
              alert(result.message); // Отображаем сообщение от сервера
              form.reset(); // Сброс формы после отправки
          } else {
              alert('Ошибка при отправке данных');
          }
      } catch (error) {
          alert('Ошибка при отправке: ' + error.message);
      }
  });
});



