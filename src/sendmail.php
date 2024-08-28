<?php
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;

require 'phpmailer/src/Exception.php';
require 'phpmailer/src/PHPMailer.php';

$mail = new PHPMailer(true);
$mail->CharSet = 'UTF-8';
$mail->setLanguage('ru', 'phpmailer/language/');
$mail->isHTML(true);
$mail->setFrom('skhvalchev@gmail.com', 'Гость'); // Исправлено setForm на setFrom
$mail->addAddress('skhvalchev@mail.ru');
$mail->Subject = 'Привет, Я тоже хочу на свадьбу';

$body = ''; // Инициализация переменной $body

if (!empty(trim($_POST['surename']))) { // Исправлено условие
    $body .= '<p>Будет на свадьбе ' . htmlspecialchars($_POST['surename']) . '.</p>'; // Добавлено htmlspecialchars для безопасности
}

$mail->Body = $body;

if (!$mail->send()) {
    $message = 'Ошибка: ' . $mail->ErrorInfo; // Добавлено отображение ошибки
} else {
    $message = 'Данные отправлены!';
}

$response = ['message' => $message];
header('Content-type: application/json');
echo json_encode($response);
?>
