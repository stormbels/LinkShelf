function createSuggestionElement(tagName, className, textContent) {
    const element = document.createElement(tagName);
    element.className = className;
    element.textContent = textContent;
    return element;
}


window.buildSearchableLinks = function (searchableLinkButtons) {
    return Array.from(searchableLinkButtons).map(function (button) {
        const title = button.dataset.title || "";
        const url = button.dataset.url || "";
        const tags = button.dataset.tags || "";

        return {
            id: button.dataset.id,
            title: title,
            url: url,
            tags: tags,
            searchText: `${title} ${url} ${tags}`.toLowerCase(),
        };
    });
};

window.renderEmptySearchSuggestion = function (searchSuggestions) {
    const emptySuggestion = createSuggestionElement(
        "div",
        "search-suggestion-empty",
        "No treasures found",
    );

    searchSuggestions.appendChild(emptySuggestion);
    searchSuggestions.hidden = false;
};

window.createSearchSuggestionButton = function (link, index, onSelect, onHover) {
    const suggestionButton = createSuggestionElement("button", "search-suggestion-item", "");
    suggestionButton.type = "button";

    const titleElement = createSuggestionElement("span", "search-suggestion-title", link.title || link.url);
    const urlElement = createSuggestionElement("span", "search-suggestion-url", link.url);

    suggestionButton.appendChild(titleElement);
    suggestionButton.appendChild(urlElement);

    if (link.tags !== "") {
        const tagsElement = createSuggestionElement("span", "search-suggestion-tags", link.tags);
        suggestionButton.appendChild(tagsElement);
    }
    suggestionButton.addEventListener("click", function () {
        onSelect(link);
    });

    suggestionButton.addEventListener("mouseenter", function () {
        onHover(index);
    });

    return suggestionButton;
};

window.updateActiveSearchSuggestion = function (searchSuggestions, activeSuggestionIndex) {
    const suggestionButtons = searchSuggestions.querySelectorAll(".search-suggestion-item");

    suggestionButtons.forEach(function (button, index) {
        button.classList.toggle("is-active", index === activeSuggestionIndex);
    });
};

window.nextSearchSuggestionIndex = function (currentIndex, suggestionsCount, step) {
    if (suggestionsCount === 0) {
        return -1;
    }

    return (currentIndex + step + suggestionsCount) % suggestionsCount;
};
