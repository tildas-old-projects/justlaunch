function init () {
  if (!checkJava()) {
    document.getElementById('java-not-found').classList.add('p-centered')
  }
}
