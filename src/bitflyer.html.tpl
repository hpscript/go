<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>BTC / ETH / XRP / XML / MONA</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
</head>
<body>
	<section class="section">
    <div class="container">
      <h1 class="title">
        BTC / ETH / XRP / XML / MONA
      </h1>
      <p class="subtitle">
        ビットコイン取引の<strong>現在値</strong>が変わる!
      </p>
      <div class="tabs">
      	<ul id="mylist">
      		<li><a href="/btc">BTC</a></li>
      		<li><a href="/eth">ETH</a></li>
      		<li><a href="/xrp">EXP</a></li>
      		<li><a href="/xlm">XLM</a></li>
      		<li><a href="/mona">MONA</a></li>
      	</ul>
      </div>
      <div class="box">
 	  {{ . }}円
	  </div>
      <div class="notification is-warning is-light">
  		※BTCはビットコイン、ETHはイーサリアム、EXPはリップル、XLMはステラルーメン、MONAはモナコインです。
	  </div>
    </div>
  </section>
  <script>
  	const pathname = location.pathname;
  	var i;
  	switch(pathname){
  		case '/btc':
  			i = 0;break;
  		case '/eth':
  			i = 1;break;
  		case '/xrp':
  			i = 2;break;
  		case '/xlm':
  			i = 3;break;
  		case '/mona':
  			i = 4;break;
  	}
  	var cols = document.querySelectorAll('#mylist li');
  	cols[i].classList.add('is-active');
  </script>
</body>
</html>