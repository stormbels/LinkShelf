window.pageParams = new URLSearchParams(window.location.search);

window.currentPageURLWithoutStateParams = function () {
    const cleanURL = new URL(window.location.href);

    [
        "error",
        "success",
        "id",
        "url",
        "title",
        "tags",
        "existing_id",
        "existing_url",
        "existing_title",
        "existing_tags",
    ].forEach(function (param) {
        cleanURL.searchParams.delete(param);
    });

    return cleanURL.pathname + cleanURL.search;
};

window.removeStateParamsFromAddressBar = function () {
    const cleanURL = window.currentPageURLWithoutStateParams();
    window.history.replaceState({}, "", cleanURL || "/search");
};
