const quickLinkForm = document.querySelector("#quickLinkForm");
const quickURLInput = document.querySelector("#url");
const linkModal = document.querySelector("#linkModal");
const linkModalForm = document.querySelector("#linkModal .link-modal-form");
const modalURLInput = document.querySelector("#modal-url");
const modalTitleInput = document.querySelector("#modal-title");
const modalTagsInput = document.querySelector("#modal-tags");
const modalCloseButton = document.querySelector(".modal-close");
const createSubmitButton = linkModalForm.querySelector('button[type="submit"]');
const modalActions = linkModal.querySelector(".modal-actions, .form-actions, .state-actions");

window.enableOrderedArrowNavigation(linkModal, [
    modalURLInput,
    modalTitleInput,
    modalTagsInput,
    createSubmitButton,
]);

window.enableHorizontalButtonNavigation(modalActions);

quickLinkForm.addEventListener("submit", function (event) {
    event.preventDefault();

    modalURLInput.value = quickURLInput.value;
    linkModal.showModal();
    modalTitleInput.focus();
});


modalCloseButton.addEventListener("click", function () {
    linkModal.close();
});
