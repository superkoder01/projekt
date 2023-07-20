import { defineStore } from 'pinia';

export const useThemeStore = defineStore({
  id: 'themeStore',
  state: () => {
    return {
      theme: '',
      isDark: '',
      logoURL: ''
    };
  },
  actions: {
    setTheme (theme: string) {
      this.theme = theme;
    },
    setIsDark (isDark: string) {
      this.isDark = isDark;
    },
    setLogoURL (url: string) {
      this.logoURL = url;
    }
  },
  getters: {
    getTheme (): string {
      return this.theme;
    },
    getIsDark (): string {
      return this.isDark;
    },
    getLogoURL (): string {
      return this.logoURL;
    }
  }
});
