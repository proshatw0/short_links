const express = require('express');
const app = express();
const net = require('net');
const path = require('path');

app.use(express.static('public'));

app.use((req, res, next) => {
  if (req.path === '/') {
    res.redirect('/create_link');
  } else {
    next();
  }
});

app.post('/create_link', (req, res) => {
  let inputValue = ``;

  req.on('data', (chunk) => {
    inputValue += chunk;
  });

  req.on('end', () => {
    const client = new net.Socket();

    client.connect(6379, '10.241.125.222', () => {
      const request = `post\n${inputValue}\n"`;
      client.write(request);
    });

    let responseData = '';

    client.on('data', (data) => {
      responseData += data.toString();
    });

    client.on('end', () => {
      // Вернуть результат клиенту
      res.end(responseData);
      client.end();
    });

    client.on('error', (error) => {
      console.error(`Ошибка при подключении к серверу: ${error}`);
      res.status(500).end('Internal Server Error');
    });
  });
});

app.get('/:value', (req, res) => {
  const value = req.params.value;
  if (value === "favicon.ico") {
    res.status(204).end();
    return; 
  }
  if (value === "create_link"){
    const filePath = path.resolve("site", 'index.html');
    res.sendFile(filePath);
    return; 
  }

  const client = new net.Socket();
  client.connect(6379, '10.241.125.222', () => {
    const request = `get\nhttp://10.241.125.222/${value}\n`;
    client.write(request);
  });

  let responseData = ''; 

  client.on('data', (data) => {
    responseData += data.toString();
  });

  client.on('end', () => {
    res.redirect(responseData);

    client.end();
  });

  client.on('error', (error) => {
    console.error(`Ошибка при подключении к серверу: ${error}`);
    res.status(500).send('Internal Server Error');
  });
});

app.listen(80, '10.241.125.222', () => {
  console.log('Сервер запущен на порту 80');
});