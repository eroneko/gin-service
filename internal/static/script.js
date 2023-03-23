let form = document.querySelector('form');

form.addEventListener('submit', (e) => {
  e.preventDefault();
  const username = form.elements.username.value;
  const password = form.elements.password.value;

  // 发送请求并处理数据
  fetch('/sessions', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({username: username, password: password})
  })
      .then(response => response.json())
      .then(data => {
        // 获取令牌
        const token = data.token;

        // 发送第二个请求，将令牌放入头部中
        fetch('/user/info', {
          headers: {
            'Authorization': 'Bearer ' + token
          }
        })
            .then(response => response.json())
            .then(data => {
              // 处理数据
              console.log(data);
              const dataString = JSON.stringify(data);
              window.location.href = `info?data=${dataString}`;
            })
            .catch(error => {
              // 处理错误
              console.error(error);
            });
      })
      .catch(error => {
        // 处理错误
        console.error(error);
      });
});