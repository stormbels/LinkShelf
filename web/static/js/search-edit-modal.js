window.openSearchEditModal = function (linkID, linkURL, linkTitle, linkTags) {
    const redirectURL = window.currentPageURLWithoutStateParams();

    window.openEditLinkModal({
        id: linkID,
        url: linkURL || "",
        title: linkTitle || "",
        tags: linkTags || "",
        action: `/links/${linkID}/update?target=${encodeURIComponent(redirectURL)}`,
    });
};

window.openSearchEditModalByURL = function (linkURL) {
    const normalizedLinkURL = window.normalizeURLForCompare(linkURL);
    const matchingEditButton = Array.from(document.querySelectorAll(".edit-button")).find(function (button) {
        return window.normalizeURLForCompare(button.dataset.url) === normalizedLinkURL;
    });

    if (!matchingEditButton) {
        return false;
    }

    window.openSearchEditModal(
        matchingEditButton.dataset.id,
        matchingEditButton.dataset.url,
        matchingEditButton.dataset.title,
        matchingEditButton.dataset.tags,
    );

    return true;
};

window.openSearchEditModalFromParams = function (prefix) {
    const linkData = window.readEditLinkParams(window.pageParams, prefix);

    if (!linkData.id) {
        return;
    }

    window.openSearchEditModal(linkData.id, linkData.url, linkData.title, linkData.tags);
};
