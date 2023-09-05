function getVersion() {
    var xhr = new XMLHttpRequest();
    var versionFooter = document.getElementById("versionFooter");
    var githubRepository = "https://github.com/Sch8ill/GSearch-web"

    xhr.open("GET", "/api/version");
    xhr.onload = function () {
        if (xhr.status === 200) {
            versionFooter.innerHTML = '<a href="' + githubRepository + '" target="_blank">' + xhr.responseText + '</a>';
        }
    };
    xhr.send();
}
getVersion()