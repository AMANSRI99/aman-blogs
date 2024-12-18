window.onload = function() {
    const storageKey = 'theme-preference';

    function getColorPreference() {
        if (localStorage.getItem(storageKey))
            return localStorage.getItem(storageKey);
        return window.matchMedia('(prefers-color-scheme: dark)').matches
            ? 'dark'
            : 'light';
    }

    function reflectPreference() {
        document.firstElementChild.setAttribute('data-theme', theme.value);
        document.querySelector('#theme-toggle')?.setAttribute('aria-label', theme.value);

        // Toggle dark class on html element
        if (theme.value === 'light') {
            document.documentElement.classList.add('dark');
            document.body.classList.remove('bg-slate-900');
            document.body.classList.add('bg-white');
        } else {
            document.documentElement.classList.remove('dark');
            document.body.classList.add('bg-slate-900');
            document.body.classList.remove('bg-white');
        }
    }

    function setPreference() {
        localStorage.setItem(storageKey, theme.value);
        reflectPreference();
    }

    const theme = {
        value: getColorPreference(),
    }

    // set initial preference
    reflectPreference();

    document.querySelector('#theme-toggle').addEventListener('click', () => {
        theme.value = theme.value === 'light' ? 'dark' : 'light';
        setPreference();
    });

    // sync with system changes
    window.matchMedia('(prefers-color-scheme: dark)')
        .addEventListener('change', ({matches:isDark}) => {
            theme.value = isDark ? 'dark' : 'light';
            setPreference();
        });
}