const successDialog = document.querySelector("#successDialog");
const successCloseButton = document.querySelector(".success-close-button");

const errorDialogs = document.querySelectorAll(".state-dialog-error");
const stateCloseButtons = document.querySelectorAll(".state-close-button");
const searchDuplicateEditButton = document.querySelector(".search-duplicate-edit-button");
const searchInvalidEditButton = document.querySelector(".search-invalid-edit-button");
const searchBrokenEditButton = document.querySelector(".search-broken-edit-button");

errorDialogs.forEach(function (dialog) {
    dialog.showModal();

    const actions = dialog.querySelector(".state-actions");
    if (actions) {
        enableHorizontalButtonNavigation(actions);
    }

    window.enableButtonArrowNavigation(dialog);

    const firstButton = dialog.querySelector("button");
    if (firstButton) {
        firstButton.focus();
    }
});

stateCloseButtons.forEach(function (button) {
    button.addEventListener("click", function () {
        const dialog = button.closest("dialog");
        if (dialog) {
            dialog.close();
        }

        removeStateParamsFromAddressBar();
    });
});

if (searchDuplicateEditButton) {
    searchDuplicateEditButton.addEventListener("click", function () {
        const dialog = searchDuplicateEditButton.closest("dialog");
        if (dialog) {
            dialog.close();
        }

        const duplicateURL = pageParams.get("existing_url") || pageParams.get("url");
        const openedFromVisibleResult = duplicateURL && openSearchEditModalByURL(duplicateURL);
        if (!openedFromVisibleResult) {
            openSearchEditModalFromParams("existing_");
        }

        removeStateParamsFromAddressBar();
    });
}

function openEditModalFromStateButton(button) {
    const dialog = button.closest("dialog");
    if (dialog) {
        dialog.close();
    }

    openSearchEditModalFromParams("");
    removeStateParamsFromAddressBar();
}

searchInvalidEditButton?.addEventListener("click", function () {
    openEditModalFromStateButton(searchInvalidEditButton);
});

searchBrokenEditButton?.addEventListener("click", function () {
    openEditModalFromStateButton(searchBrokenEditButton);
});

if (successDialog) {
    successDialog.showModal();
    window.enableButtonArrowNavigation(successDialog);
    successCloseButton?.focus();
}

if (successCloseButton) {
    successCloseButton.addEventListener("click", function () {
        successDialog.close();
        removeStateParamsFromAddressBar();
    });
}
