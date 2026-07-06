const stateDialog = document.querySelector("#stateDialog");
const successDialog = document.querySelector("#successDialog");
const successCloseButton = document.querySelector(".success-close-button");
const stateCloseButtons = document.querySelectorAll(".state-close-button");

function enableDialogButtonNavigation(dialog) {
    if (typeof window.enableButtonArrowNavigation === "function") {
        window.enableButtonArrowNavigation(dialog);
    }
}

function showDialog(dialog) {
    if (!dialog) {
        return;
    }

    dialog.showModal();
    enableDialogButtonNavigation(dialog);
    dialog.querySelector("button")?.focus();
}

function closeDialogAndClearURL(dialog) {
    dialog?.close();
    window.history.replaceState({}, "", "/");
}

successCloseButton?.addEventListener("click", function () {
    closeDialogAndClearURL(successDialog);
});

stateCloseButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        closeDialogAndClearURL(stateDialog);
    });
});

showDialog(stateDialog);
showDialog(successDialog);
