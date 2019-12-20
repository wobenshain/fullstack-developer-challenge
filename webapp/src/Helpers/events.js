export const setStateOnChange = (set) => ({ target }) => {
    const { value } = target;
    set(value);
};
