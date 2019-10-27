function init () {
  if (!checkJava()) {
    document.getElementById('java-not-found').style('display:block')
  }
}
