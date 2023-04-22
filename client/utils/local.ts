export const localStore = {
    get: (key: string) => {
        let result = localStorage.getItem(key) || '{}' || '[]' || '""'
        return JSON.parse(result)
    },

    set: (key: string, value: any) => {
        if (typeof value === 'object') {
            value = JSON.stringify(value)
            localStorage.setItem(key, value)
        }
    },
}
