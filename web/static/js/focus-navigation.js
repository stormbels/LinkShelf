window.visibleFocusableElements = function (elements) {
    return Array.from(elements).filter(function (element) {
        return element && !element.disabled && element.offsetParent !== null;
    });
};

window.focusElementByStep = function (elements, step) {
    const focusableElements = window.visibleFocusableElements(elements);
    if (focusableElements.length === 0) {
        return;
    }

    const currentIndex = focusableElements.indexOf(document.activeElement);
    if (currentIndex === -1) {
        focusableElements[0].focus();
        return;
    }

    const nextIndex = (currentIndex + step + focusableElements.length) % focusableElements.length;
    focusableElements[nextIndex].focus();
};

function focusableButtons(container) {
    return window.visibleFocusableElements(container.querySelectorAll("button"));
}

function enableArrowNavigation(target, previousKeys, nextKeys, getElements) {
    if (!target) {
        return;
    }

    target.addEventListener("keydown", function (event) {
        const isPreviousKey = previousKeys.includes(event.key);
        const isNextKey = nextKeys.includes(event.key);
        if (!isPreviousKey && !isNextKey) {
            return;
        }

        const elements = getElements();
        if (elements.length === 0) {
            return;
        }

        event.preventDefault();
        const step = isNextKey ? 1 : -1;
        window.focusElementByStep(elements, step);
    });
}

window.enableButtonArrowNavigation = function (dialog) {
    enableArrowNavigation(
        dialog,
        ["ArrowLeft", "ArrowUp"],
        ["ArrowRight", "ArrowDown"],
        function () {
            return focusableButtons(dialog);
        },
    );
};

window.enableOrderedArrowNavigation = function (dialog, orderedElements) {
    enableArrowNavigation(
        dialog,
        ["ArrowUp"],
        ["ArrowDown"],
        function () {
            return orderedElements;
        },
    );
};


window.enableHorizontalButtonNavigation = function (container) {
    enableArrowNavigation(
        container,
        ["ArrowLeft"],
        ["ArrowRight"],
        function () {
            return focusableButtons(container);
        },
    );
};
