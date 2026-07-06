const editButtons = document.querySelectorAll(".edit-button");
const editLinkModal = document.querySelector("#editLinkModal");
const editLinkForm = document.querySelector("#editLinkForm");
const editModalURLInput = document.querySelector("#edit-modal-url");
const editModalTitleInput = document.querySelector("#edit-modal-title");
const editModalTagsInput = document.querySelector("#edit-modal-tags");
const editModalCloseButtons = document.querySelectorAll(".edit-modal-close");

const editSubmitButton = editLinkForm.querySelector('button[type="submit"]');
const editModalActions = editLinkModal.querySelector(".modal-actions, .form-actions, .state-actions");

window.enableOrderedArrowNavigation(editLinkModal, [
    editModalURLInput,
    editModalTitleInput,
    editModalTagsInput,
    ...Array.from(editModalActions?.querySelectorAll("button") || []),
]);

window.enableHorizontalButtonNavigation(editModalActions);

function updateEditSubmitState() {
    const hasChanges = editModalURLInput.value !== editLinkForm.dataset.originalUrl ||
        editModalTitleInput.value !== editLinkForm.dataset.originalTitle ||
        editModalTagsInput.value !== editLinkForm.dataset.originalTags;

    editSubmitButton.disabled = !hasChanges;
}

window.openEditLinkModal = function (linkData) {
    editLinkForm.action = linkData.action || `/links/${linkData.id}/update`;
    editModalURLInput.value = linkData.url;
    editModalTitleInput.value = linkData.title;
    editModalTagsInput.value = linkData.tags;

    editLinkForm.dataset.originalUrl = editModalURLInput.value;
    editLinkForm.dataset.originalTitle = editModalTitleInput.value;
    editLinkForm.dataset.originalTags = editModalTagsInput.value;
    updateEditSubmitState();

    editLinkModal.showModal();
    editModalTitleInput.focus();
};

window.readEditLinkParams = function (params, prefix = "") {
    return {
        id: params.get(`${prefix}id`) || params.get("id"),
        url: params.get(`${prefix}url`) || params.get("url"),
        title: params.get(`${prefix}title`) || params.get("title"),
        tags: params.get(`${prefix}tags`) || params.get("tags"),
    };
};

window.normalizeURLForCompare = function (rawURL) {
    if (!rawURL) {
        return "";
    }

    try {
        const parsedURL = new URL(rawURL);
        parsedURL.pathname = parsedURL.pathname.replace(/\/+$/, "");
        return parsedURL.toString();
    } catch {
        return rawURL.replace(/\/+$/, "");
    }
};

editButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        window.openEditLinkModal({
            id: button.dataset.id,
            url: button.dataset.url,
            title: button.dataset.title,
            tags: button.dataset.tags,
        });
    });
});

[editModalURLInput, editModalTitleInput, editModalTagsInput].forEach(function (input) {
    input.addEventListener("input", updateEditSubmitState);
});

editModalCloseButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        editLinkModal.close();
    });
});
