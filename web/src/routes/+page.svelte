<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <link
      rel="stylesheet"
      href="https://unpkg.com/papercss@1.9.2/dist/paper.min.css"
    />
    <title>Minit</title>
  </head>
  <style>
    * {
      padding: 0;
      margin: 0;
    }
    body {
      padding: 1rem;
    }
    input {
      margin-bottom: 1rem;
    }
    .title {
      text-align: center;
    }
    .container {
      display: grid;
      place-content: center;
    }
    .result {
      margin-top: 1.5rem;
    }
    .footer {
      text-align: center;
      position: fixed;
      bottom: 0;
      padding-bottom: 1rem;
      width: 100%;
    }
  </style>
  <body>
    <div class="paper container">
      <h1 class="title">Minit</h1>
      <input type="url" placeholder="URL To Shorten" id="uri" />
      <button class="btn-block btn-primary" onclick="shortenURI()">
        Shorten
      </button>
      <div class="container result">
        <p><a id="key" /></p>
      </div>
    </div>
    <div class="footer">All Generated Links Expire After 24 Hours</div>
    <script>
      function shortenURI() {
        let host = "http://localhost:8080";
        let data = {
          body: document.getElementById("uri").value,
        };
        let request = new XMLHttpRequest();
        request.open("POST", host + "/shorten", true);
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify(data));
        request.onload = function () {
          if (request.status === 200) {
            document.getElementById("key").textContent =
              host + "/" + request.response;
            document.getElementById("key").href = host + "/" + request.response;
          } else {
            document.getElementById("key").removeAttribute("href");
            document.getElementById("key").textContent =
              "Please enter a valid URL";
          }
        };
      }
    </script>
  </body>
</html>
