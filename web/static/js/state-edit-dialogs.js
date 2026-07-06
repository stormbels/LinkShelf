const duplicateEditButton = document.querySelector(".duplicate-edit-button");
const invalidEditEditButton = document.querySelector(".invalid-edit-edit-button");
const brokenEditEditButton = document.querySelector(".broken-edit-edit-button");
const stateDialog = document.querySelector("#stateDialog");
const pageParams = new URLSearchParams(window.location.search);
const editButtons = document.querySelectorAll("[data-id][data-url][data-title][data-tags]");

function closeStateDialog() {
    stateDialog?.close();
    window.history.replaceState({}, "", "/");
}

function openEditModalFromParams(prefix) {
    const linkData = window.readEditLinkParams(pageParams, prefix);

    if (!linkData.id) {
        return false;
    }

    window.openEditLinkModal(linkData);
    return true;
}

function openEditModalByURL(linkURL) {
    const normalizedLinkURL = window.normalizeURLForCompare(linkURL);
    const matchingEditButton = Array.from(editButtons).find(function (button) {
        return window.normalizeURLForCompare(button.dataset.url) === normalizedLinkURL;
    });

    if (!matchingEditButton) {
        return false;
    }

    window.openEditLinkModal({
        id: matchingEditButton.dataset.id,
        url: matchingEditButton.dataset.url,
        title: matchingEditButton.dataset.title,
        tags: matchingEditButton.dataset.tags,
    });

    return true;
}

function openOriginalEditModalFromState() {
    closeStateDialog();
    openEditModalFromParams("");
}

duplicateEditButton?.addEventListener("click", function () {
    const duplicateURL = pageParams.get("existing_url") || pageParams.get("url");

    closeStateDialog();

    const openedFromVisibleLink = duplicateURL && openEditModalByURL(duplicateURL);
    if (!openedFromVisibleLink) {
        openEditModalFromParams("existing_");
    }
});

invalidEditEditButton?.addEventListener("click", openOriginalEditModalFromState);
brokenEditEditButton?.addEventListener("click", openOriginalEditModalFromState);
