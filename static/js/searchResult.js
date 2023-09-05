function getSearchQuery() {
    var urlParams = new URLSearchParams(window.location.search);
    return urlParams.get("q")
}

function setSearchInputValue(searchQuery) {
    var searchInput = document.getElementById("searchInput");
    searchInput.value = searchQuery;
}

function populateResults(searchQuery) {
    getSearchResults(searchQuery)
        .then(searchResults => {
            displayResults(searchResults.Results);
        })
}

function getSearchResults(searchQuery) {
    return fetch("/api/search?q=" + searchQuery)
        .then(response => {
            return response.json();
        });
}

function removeDuplicateResults(results) {
    var cleanResults = []
    var double = false

    for (var i = 0; i < results.length; i++) {
        var result = results[i];

        for (var j; j < cleanResults.length; j++) {
            var cleanResult = cleanResults[j]
            if (result.Url == cleanResult.Url) {
                double = true
            }
        }

        if (!double) {
            cleanResults.push(result)
        } else {
            double = false
        }
    }
    return cleanResults
}

function displayResults(results) {
    var resultList = document.getElementById("results");
    resultList.innerHTML = "";
    console.log(results)

    if (Object.is(results, null)) {
        console.log("no results...");
        var noResults = document.createElement("p");
        noResults.innerHTML = "Keine Ergebnisse gefunden ):";
        resultList.appendChild(noResults);
        return    
    }

    results = removeDuplicateResults(results);

    for (var i = 0; i < results.length; i++) {
        var searchResult = results[i];
        console.log(searchResult);

        var resultItem = document.createElement("li");
        resultItem.innerHTML = '<a href="' + searchResult.Url + '" target="_blank">' + searchResult.Url + '</a> \
         <ul> \
            <li>' + searchResult.Timestamp + '</li> \
            <li>' + searchResult.Text + '</li> \
         </ul> \
         <br> \
         <br> \
         ';
        resultList.appendChild(resultItem);
    }
}


var searchQuery = getSearchQuery();
setSearchInputValue(searchQuery);
populateResults(searchQuery);
