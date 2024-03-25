/*
Below are a full unminified and minified version of the JavaScript snippet that
you can add to the head of your html file. The snippet will reload the page when
the WebSocket connection is closed. the snippet is used in the
example/views/hello.html file. 

You can dial in the timeout value to what works
on your system.  
*/

// FULL
{
  let active_full = false;
  // set the URL of the websocket server to the host where the go app is running
  sock = new WebSocket("ws://localhost:8080/reload");
  sock.onopen = function (event) {
    console.log("connected");
    active = true;
  };

  sock.onclose = function (event) {
    console.log("disconnected");
    // the timeout value needs to be long enough for the
    // go app to reload before refreshing this page.
    // tune it to what works on your system.
    if (active_full) {
      setTimeout(function () {
        location.reload();
        active_full = false;
      }, 2000);
    }
  };
}

// MINIFIED and console.log removed
// prettier-ignore
{let o=!1;sock=new WebSocket("ws://localhost:8080/reload"),sock.onopen=function(o){active=!0},sock.onclose=function(e){o&&setTimeout((function(){location.reload(),o=!1}),2e3)}}
