const pageParams = new URLSearchParams(window.location.search);
const stateDialog = document.querySelector("#stateDialog");
const quickURLInput = document.querySelector("#url");
const linkModal = document.querySelector("#linkModal");
const modalTitleInput = document.querySelector("#modal-title");
const modalTagsInput = document.querySelector("#modal-tags");
const modalURLInput = document.querySelector("#modal-url");
const invalidCreateEditButton = document.querySelector(".invalid-create-edit-button");
const brokenCreateEditButton = document.querySelector(".broken-create-edit-button");

function restoreCreateModalFromState() {
    const params = pageParams;

    stateDialog.close();

    window.requestAnimationFrame(function () {
        if (quickURLInput) {
            quickURLInput.value = params.get("url") || "";
        }
        modalURLInput.value = params.get("url") || "";
        modalTitleInput.value = params.get("title") || "";
        modalTagsInput.value = params.get("tags") || "";

        linkModal.showModal();
        modalURLInput.focus();

        window.history.replaceState({}, "", "/");
    });
}

invalidCreateEditButton?.addEventListener("click", restoreCreateModalFromState);
brokenCreateEditButton?.addEventListener("click", restoreCreateModalFromState);
