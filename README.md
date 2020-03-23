<h2 align="center">Platform Independent Really-Cool Adventure Text Engine</h2>
<p align="center">
    <img height="200px" src="assets/pirate.png">
</p>
<p align="center">
    <i>Crazy motto goes here!</i>
</p>

<p align="center">
  <a href="https://github.com/retrogen/pirate/actions">
    <img src="https://github.com/retrogen/pirate/workflows/test/badge.svg" />
  </a>
  <a href="https://codecov.io/gh/retrogen/pirate">
    <img src="https://codecov.io/gh/retrogen/pirate/branch/master/graph/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/retrogen/pirate">
    <img src="https://goreportcard.com/badge/github.com/retrogen/pirate" />
  </a>
  <a href="https://github.com/retrogen/pirate/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/retrogen/pirate.svg">
  </a>
</p>

# About

This project exists to bring new era of interactive fiction games! 
There are number of awesome text game engines: TADS, Inform, INSTEAD, Twine and others. 
But they all have one big disadvantage: you need to install some software on your device to play IF game.
Why do you need install something to play IF game? It's ridiculous!
So we came up with idea: make IF game accessible on every device you can imagine, 
moreover we want to make it cheap. 
Is it awesome to play using browser, your favorite messenger, maybe some CLI interface or just curl if you want to go hardcore...

This project consists of two big, but very related parts: compiler for source of IF game and VM that can run it on server.
With such simple principle you can access server through HTTP, send you command to game, get response and render it to user.
Our architecture also propose new exciting genre of IF game, that wasn't explored enough: online IF games!

### License

pirate is [MIT licensed](./LICENSE).
