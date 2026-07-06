const searchForm = document.querySelector("#searchForm");
const searchInput = document.querySelector("#search");
const searchSuggestions = document.querySelector("#searchSuggestions");
const searchableLinkButtons = document.querySelectorAll(".edit-button");
const searchableLinks = window.buildSearchableLinks(searchableLinkButtons);

let activeSuggestionIndex = -1;
let currentSuggestions = [];


function navigateToSearchResult(link) {
    const params = new URLSearchParams();
    params.set("q", searchInput.value.trim());
    params.set("selected", link.id);

    window.location.href = `/search?${params.toString()}`;
}

function hideSearchSuggestions() {
    searchSuggestions.hidden = true;
    activeSuggestionIndex = -1;
}

function activateSearchSuggestion(index) {
    activeSuggestionIndex = index;
    window.updateActiveSearchSuggestion(searchSuggestions, activeSuggestionIndex);
}

function moveActiveSuggestion(step) {
    activeSuggestionIndex = window.nextSearchSuggestionIndex(
        activeSuggestionIndex,
        currentSuggestions.length,
        step,
    );
    window.updateActiveSearchSuggestion(searchSuggestions, activeSuggestionIndex);
}

function renderSuggestionButtons() {
    currentSuggestions.forEach(function (link, index) {
        const suggestionButton = window.createSearchSuggestionButton(
            link,
            index,
            navigateToSearchResult,
            activateSearchSuggestion,
        );

        searchSuggestions.appendChild(suggestionButton);
    });
}

function handleSearchInputKeydown(event) {
    if (event.key === "ArrowDown" && currentSuggestions.length > 0) {
        event.preventDefault();
        moveActiveSuggestion(1);
        return;
    }

    if (event.key === "ArrowUp" && currentSuggestions.length > 0) {
        event.preventDefault();
        moveActiveSuggestion(-1);
        return;
    }

    if (event.key === "Enter" && activeSuggestionIndex >= 0) {
        event.preventDefault();
        navigateToSearchResult(currentSuggestions[activeSuggestionIndex]);
        return;
    }

    if (event.key === "Escape") {
        hideSearchSuggestions();
    }
}

function renderSearchSuggestions() {
    const query = searchInput.value.trim().toLowerCase();
    searchSuggestions.innerHTML = "";
    activeSuggestionIndex = -1;

    if (query === "") {
        currentSuggestions = [];
        hideSearchSuggestions();
        return;
    }

    currentSuggestions = searchableLinks.filter(function (link) {
        return link.searchText.includes(query);
    }).slice(0, 6);

    if (currentSuggestions.length === 0) {
        window.renderEmptySearchSuggestion(searchSuggestions);
        return;
    }

    renderSuggestionButtons();
    searchSuggestions.hidden = false;
}

if (searchInput) {
    searchInput.addEventListener("input", renderSearchSuggestions);

    searchInput.addEventListener("keydown", handleSearchInputKeydown);
}

if (searchForm) {
    searchForm.addEventListener("submit", hideSearchSuggestions);
}

document.addEventListener("click", function (event) {
    if (searchForm && !searchForm.contains(event.target)) {
        hideSearchSuggestions();
    }
});
