var button = document.getElementById('button')

// 、、alter(button);
button.onclick = function() {
    // alter("123");
    xhr = new XMLHttpRequest();
    xhr.open('GET', 'https://os.ncuos.com/api/user/token');
    xhr.send();

    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            alert(xhr.response);
        } else {
            alert("账户或密码错误");
        }
    }
}