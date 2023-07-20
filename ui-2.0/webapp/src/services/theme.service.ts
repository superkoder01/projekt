import { useUserStore } from './../store/user.store';
import {ServiceTypeEnum} from "@/services/logger/service-type.enum";
import { ThemeInfo } from "@/models/theme-info";
import { useThemeStore } from "@/store/theme.store";

const logo = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjIxIiBoZWlnaHQ9IjEwMyIgdmlld0JveD0iMCAwIDIyMSAxMDMiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+DQo8ZyBjbGlwLXBhdGg9InVybCgjY2xpcDBfODI2XzE0NTc2KSI+DQo8cGF0aCBkPSJNMTA4LjAyOSA5Ni43OTQzQzk4LjQ1IDk2Ljc3MDQgODkuMDkyNyA5My45MTE3IDgxLjEzOTUgODguNTc5M0M3My4xODYzIDgzLjI0NjkgNjYuOTkzOSA3NS42OCA2My4zNDQ2IDY2LjgzNDVDNTkuNjk1MyA1Ny45ODkgNTguNzUyOSA0OC4yNjE4IDYwLjYzNjQgMzguODgxNkM2Mi41MTk4IDI5LjUwMTUgNjcuMTQ0NiAyMC44ODkzIDczLjkyNjYgMTQuMTMyOUM4MC43MDg2IDcuMzc2NDkgODkuMzQzNSAyLjc3OTA2IDk4Ljc0MDQgMC45MjE0MDhDMTA4LjEzNyAtMC45MzYyNDIgMTE3Ljg3NSAwLjAyOTIxNTkgMTI2LjcyMyAzLjY5NThDMTM1LjU3IDcuMzYyMzggMTQzLjEzMiAxMy41NjU2IDE0OC40NTEgMjEuNTIxOEMxNTMuNzcxIDI5LjQ3OCAxNTYuNjEgMzguODMwMyAxNTYuNjEgNDguMzk3MkMxNTYuNTc4IDYxLjI0NCAxNTEuNDQ1IDczLjU1MzcgMTQyLjMzOCA4Mi42MjY0QzEzMy4yMzEgOTEuNjk5MiAxMjAuODkzIDk2Ljc5NDMgMTA4LjAyOSA5Ni43OTQzWk0xMDguMDI5IDMuNzAxMzJDOTkuMTgyNCAzLjcyMjUzIDkwLjU0MDMgNi4zNjE4OCA4My4xOTQ4IDExLjI4NkM3NS44NDkyIDE2LjIxMDEgNzAuMTI5OCAyMy4xOTgxIDY2Ljc1ODkgMzEuMzY3QzYzLjM4ODEgMzkuNTM2IDYyLjUxNzEgNDguNTE5NCA2NC4yNTYxIDU3LjE4MjVDNjUuOTk1IDY1Ljg0NTYgNzAuMjY1OCA3My43OTk2IDc2LjUyOSA4MC4wMzk3QzgyLjc5MjEgODYuMjc5OCA5MC43NjY3IDkwLjUyNiA5OS40NDUxIDkyLjI0MkMxMDguMTI0IDkzLjk1NzkgMTE3LjExNiA5My4wNjY1IDEyNS4yODggODkuNjgwNUMxMzMuNDU5IDg2LjI5NDQgMTQwLjQ0MyA4MC41NjU2IDE0NS4zNTUgNzMuMjE3N0MxNTAuMjY4IDY1Ljg2OTkgMTUyLjg5IDU3LjIzMjYgMTUyLjg5IDQ4LjM5NzJDMTUyLjg2MiAzNi41MzMyIDE0OC4xMjMgMjUuMTY0OSAxMzkuNzEzIDE2Ljc4NTlDMTMxLjMwMyA4LjQwNjg0IDExOS45MDkgMy43MDEyOSAxMDguMDI5IDMuNzAxMzJaIiBmaWxsPSIjMkFGRDg4Ii8+DQo8cGF0aCBkPSJNMTA4LjAyOSAzLjIwMTMyTDEwOC4wMjggMy4yMDEzMkM5OS4wODI1IDMuMjIyNzcgOTAuMzQ0IDUuODkxNTYgODIuOTE2MyAxMC44NzA3Qzc1LjQ4ODcgMTUuODQ5OCA2OS43MDUzIDIyLjkxNTkgNjYuMjk2NyAzMS4xNzYzQzYyLjg4ODEgMzkuNDM2NyA2Mi4wMDc0IDQ4LjUyMDggNjMuNzY1OCA1Ny4yODA5QzY1LjUyNDMgNjYuMDQxIDY5Ljg0MjkgNzQuMDg0IDc2LjE3NjEgODAuMzkzOUM4Mi41MDkzIDg2LjcwMzggOTAuNTcyOSA5MC45OTc0IDk5LjM0ODEgOTIuNzMyNUMxMDguMTIzIDk0LjQ2NzUgMTE3LjIxNyA5My41NjYyIDEyNS40NzkgOTAuMTQyNEMxMzMuNzQyIDg2LjcxODUgMTQwLjgwMyA4MC45MjU3IDE0NS43NzEgNzMuNDk1NkMxNTAuNzM5IDY2LjA2NTYgMTUzLjM5IDU3LjMzMTYgMTUzLjM5IDQ4LjM5NzJWNDguMzk2QzE1My4zNjIgMzYuMzk5NCAxNDguNTcgMjQuOTA0MiAxNDAuMDY2IDE2LjQzMTdDMTMxLjU2MiA3Ljk1OTIgMTIwLjA0MSAzLjIwMTI5IDEwOC4wMjkgMy4yMDEzMlpNMTA4LjAzMSA5Ni4yOTQzQzk4LjU1IDk2LjI3MDcgODkuMjg5MSA5My40NDE0IDgxLjQxNzkgODguMTY0QzczLjU0NjggODIuODg2NiA2Ny40MTg0IDc1LjM5NzggNjMuODA2OCA2Ni42NDM4QzYwLjE5NTMgNTcuODg5OCA1OS4yNjI2IDQ4LjI2MzIgNjEuMTI2NiAzOC45OEM2Mi45OTA2IDI5LjY5NjkgNjcuNTY3NiAyMS4xNzM3IDc0LjI3OTUgMTQuNDg3MUM4MC45OTE1IDcuODAwNDYgODkuNTM3MyAzLjI1MDQyIDk4LjgzNzQgMS40MTE5MkMxMDguMTM4IC0wLjQyNjU5MyAxMTcuNzc1IDAuNTI4OTIyIDEyNi41MzEgNC4xNTc3QzEzNS4yODggNy43ODY0OCAxNDIuNzcxIDEzLjkyNTcgMTQ4LjAzNSAyMS43OTk3QzE1My4zIDI5LjY3MzQgMTU2LjEwOSAzOC45Mjg1IDE1Ni4xMSA0OC4zOTZDMTU2LjA3OCA2MS4xMTAxIDE1MC45OTkgNzMuMjkyOSAxNDEuOTg1IDgyLjI3MjJDMTMyLjk3MiA5MS4yNTEzIDEyMC43NjEgOTYuMjk0IDEwOC4wMzEgOTYuMjk0M1oiIHN0cm9rZT0iYmxhY2siIHN0cm9rZS1vcGFjaXR5PSIwLjIiLz4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNODIuNTM1NiAxNi4xNjI1QzgyLjUzNTYgMTcuNzUzOCA4Mi4wNjc0IDE5LjMwOTQgODEuMTkwMSAyMC42MzI1QzgwLjMxMjkgMjEuOTU1NyA3OS4wNjYgMjIuOTg2OSA3Ny42MDcyIDIzLjU5NTlDNzYuMTQ4NCAyNC4yMDQ5IDc0LjU0MzEgMjQuMzY0MiA3Mi45OTQ1IDI0LjA1MzdDNzEuNDQ1OCAyMy43NDMzIDcwLjAyMzIgMjIuOTc3IDY4LjkwNjcgMjEuODUxN0M2Ny43OTAyIDIwLjcyNjUgNjcuMDI5OCAxOS4yOTI5IDY2LjcyMTggMTcuNzMyMkM2Ni40MTM3IDE2LjE3MTQgNjYuNTcxOCAxNC41NTM3IDY3LjE3NjEgMTMuMDgzNUM2Ny43ODAzIDExLjYxMzMgNjguODAzNiAxMC4zNTY3IDcwLjExNjUgOS40NzI2NUM3MS40Mjk0IDguNTg4NTYgNzIuOTczIDguMTE2NyA3NC41NTIgOC4xMTY3Qzc2LjY2OTQgOC4xMTY3IDc4LjcgOC45NjQzOCA4MC4xOTczIDEwLjQ3MzNDODEuNjk0NSAxMS45ODIxIDgyLjUzNTYgMTQuMDI4NiA4Mi41MzU2IDE2LjE2MjVaIiBmaWxsPSIjMkFGRDg4Ii8+DQo8cGF0aCBmaWxsLXJ1bGU9ImV2ZW5vZGQiIGNsaXAtcnVsZT0iZXZlbm9kZCIgZD0iTTEwOS40MDIgOTQuOTI3NEMxMDkuNDAyIDk2LjUxODcgMTA4LjkzNCA5OC4wNzQzIDEwOC4wNTYgOTkuMzk3NEMxMDcuMTc5IDEwMC43MjEgMTA1LjkzMiAxMDEuNzUyIDEwNC40NzMgMTAyLjM2MUMxMDMuMDE1IDEwMi45NyAxMDEuNDA5IDEwMy4xMjkgOTkuODYwNyAxMDIuODE5Qzk4LjMxMiAxMDIuNTA4IDk2Ljg4OTQgMTAxLjc0MiA5NS43NzI5IDEwMC42MTdDOTQuNjU2NCA5OS40OTE0IDkzLjg5NiA5OC4wNTc4IDkzLjU4OCA5Ni40OTcxQzkzLjI3OTkgOTQuOTM2MyA5My40MzggOTMuMzE4NiA5NC4wNDIzIDkxLjg0ODRDOTQuNjQ2NSA5MC4zNzgyIDk1LjY2OTggODkuMTIxNiA5Ni45ODI3IDg4LjIzNzVDOTguMjk1NiA4Ny4zNTM1IDk5LjgzOTIgODYuODgxNiAxMDEuNDE4IDg2Ljg4MTZDMTAzLjUzNiA4Ni44ODE2IDEwNS41NjYgODcuNzI5MyAxMDcuMDYzIDg5LjIzODJDMTA4LjU2MSA5MC43NDcgMTA5LjQwMiA5Mi43OTM1IDEwOS40MDIgOTQuOTI3NFoiIGZpbGw9IiMyQUZEODgiLz4NCjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNMTU4LjQ2MSAzMS45ODIxQzE1OC40NjEgMzMuNTczNCAxNTcuOTkzIDM1LjEyOSAxNTcuMTE2IDM2LjQ1MjFDMTU2LjIzOSAzNy43NzUyIDE1NC45OTIgMzguODA2NSAxNTMuNTMzIDM5LjQxNTVDMTUyLjA3NCA0MC4wMjQ0IDE1MC40NjkgNDAuMTgzOCAxNDguOTIgMzkuODczM0MxNDcuMzcyIDM5LjU2MjkgMTQ1Ljk0OSAzOC43OTY2IDE0NC44MzIgMzcuNjcxNEMxNDMuNzE2IDM2LjU0NjEgMTQyLjk1NiAzNS4xMTI1IDE0Mi42NDggMzMuNTUxOEMxNDIuMzM5IDMxLjk5MSAxNDIuNDk4IDMwLjM3MzMgMTQzLjEwMiAyOC45MDMxQzE0My43MDYgMjcuNDMyOSAxNDQuNzI5IDI2LjE3NjMgMTQ2LjA0MiAyNS4yOTIyQzE0Ny4zNTUgMjQuNDA4MSAxNDguODk5IDIzLjkzNjMgMTUwLjQ3OCAyMy45MzYzQzE1Mi41OTUgMjMuOTM2MyAxNTQuNjI2IDI0Ljc4NCAxNTYuMTIzIDI2LjI5MjhDMTU3LjYyIDI3LjgwMTcgMTU4LjQ2MSAyOS44NDgyIDE1OC40NjEgMzEuOTgyMVoiIGZpbGw9IiMyQUZEODgiLz4NCjxwYXRoIGQ9Ik0xMjAuNzM5IDc1LjIxNTFDMTI0LjA0NSA3My41NTIgMTI3LjAxIDcxLjI3NzYgMTI5LjQ4MSA2OC41MTAyVjI4Ljg3MTFDMTI3LjAxMyAyNi4xMDAyIDEyNC4wNDcgMjMuODI1MiAxMjAuNzM5IDIyLjE2NjNWNDYuNzE5NEg5NC4yMjA0VjIyLjM2NzRDOTAuOTMxOSAyNC4wNzggODcuOTg5NSAyNi4zOTI5IDg1LjU0NDkgMjkuMTkyOVY1MS4xNTgxQzg1LjU0NTkgNTEuNzQyIDg1LjY1ODkgNTIuMzIwMyA4NS44Nzc2IDUyLjg2MTFDODYuMDk0NyA1My4zODM1IDg2LjQwNiA1My44NjA5IDg2Ljc5NTcgNTQuMjY5MUM4Ny4xNzk5IDU0LjY1NjIgODcuNjMwOCA1NC45Njk3IDg4LjEyNjMgNTUuMTk0NEM4OC42NzUyIDU1LjQzODYgODkuMjY5NCA1NS41NjIxIDg5Ljg2OTQgNTUuNTU2NEgxMjAuNjg2TDEyMC43MzkgNzUuMjE1MVoiIGZpbGw9IiMyQUZEODgiLz4NCjxwYXRoIGQ9Ik0yMTcuNTY3IDY1LjY2NzVIMTg2LjgzQzE4Ni4zMSA2NS43MjIyIDE4NS43ODQgNjUuNjU5MyAxODUuMjkxIDY1LjQ4MzJDMTg0Ljc5OCA2NS4zMDcxIDE4NC4zNTEgNjUuMDIyNCAxODMuOTgxIDY0LjY0OTdDMTgzLjYxMSA2NC4yNzcxIDE4My4zMjkgNjMuODI1OSAxODMuMTU0IDYzLjMyOTJDMTgyLjk3OSA2Mi44MzI1IDE4Mi45MTcgNjIuMzAyOCAxODIuOTcxIDYxLjc3ODdWMzUuODA0QzE4Mi45MzEgMzUuMjY4OCAxODIuOTk3IDM0LjczMDkgMTgzLjE2NiAzNC4yMjJDMTgzLjMzNSAzMy43MTMgMTgzLjYwNCAzMy4yNDMzIDE4My45NTYgMzIuODQwNUMxODQuNzQ2IDMyLjE0MDkgMTg1Ljc3OSAzMS43ODQyIDE4Ni44MyAzMS44NDgySDIxNy40MzRDMjE4LjI5MiAyOC41NTY0IDIxOS40MDQgMjUuMzM3MiAyMjAuNzYgMjIuMjJIMTg2LjgzQzE4NS4zMzggMjIuMjAwNCAxODMuODUzIDIyLjQxMyAxODIuNDI1IDIyLjg1MDNDMTgxLjIgMjMuMjI3MyAxODAuMDQxIDIzLjc5NzUgMTc4Ljk5MiAyNC41Mzk4QzE3OC4wMjQgMjUuMjA5MSAxNzcuMTYzIDI2LjAyMyAxNzYuNDM4IDI2Ljk1MzZDMTc1Ljc0NyAyNy44MzUgMTc1LjE2NiAyOC43OTg4IDE3NC43MDggMjkuODIzM0MxNzQuMjc5IDMwLjgwMDYgMTczLjk1MyAzMS44MjEgMTczLjczNyAzMi44NjczQzE3My41MzcgMzMuODA2MSAxNzMuNDM0IDM0Ljc2MzMgMTczLjQzMSAzNS43MjM1VjYxLjgwNTVDMTczLjQwNiA2My4zMTc3IDE3My42MTcgNjQuODI0NCAxNzQuMDU2IDY2LjI3MDlDMTc0LjQzMSA2Ny41MDc0IDE3NC45OTIgNjguNjc4NyAxNzUuNzE5IDY5Ljc0NEMxNzYuMzg4IDcwLjcxNjIgMTc3LjIgNzEuNTc5OCAxNzguMTI4IDcyLjMwNTJDMTc5LjAwNSA3Mi45ODk5IDE3OS45NjEgNzMuNTY2MSAxODAuOTc1IDc0LjAyMTdDMTgxLjk0IDc0LjQ1NDIgMTgyLjk0OCA3NC43ODIzIDE4My45ODIgNzUuMDAwNkMxODQuOTE4IDc1LjIwNjMgMTg1Ljg3MiA3NS4zMTQyIDE4Ni44MyA3NS4zMjI0SDIyMS4wNEMyMTkuNjM0IDcyLjIwMDMgMjE4LjQ3MyA2OC45NzIyIDIxNy41NjcgNjUuNjY3NVoiIGZpbGw9IndoaXRlIi8+DQo8cGF0aCBkPSJNMjE2Ljg3NSA1My41OTg2SDE4Ni43MzdWNDMuOTAzM0gyMTYuODc1VjUzLjU5ODZaIiBmaWxsPSJ3aGl0ZSIvPg0KPHBhdGggZD0iTTQ0LjQ1NTUgNjUuNjY3NEgxMy4zOTkyQzEyLjg4MDIgNjUuNzIgMTIuMzU2IDY1LjY1NTQgMTEuODY0OCA2NS40Nzg0QzExLjM3MzYgNjUuMzAxMyAxMC45Mjc3IDY1LjAxNjMgMTAuNTU5NSA2NC42NDM5QzEwLjE5MTIgNjQuMjcxNSA5LjkwOTkzIDYzLjgyMTIgOS43MzU5NiA2My4zMjU1QzkuNTYyIDYyLjgyOTkgOS40OTk3NyA2Mi4zMDE0IDkuNTUzNzYgNjEuNzc4NVYzNS44MDM5QzkuNTEzNDIgMzUuMjY4NyA5LjU3OTc1IDM0LjczMDggOS43NDg4NiAzNC4yMjE4QzkuOTE3OTYgMzMuNzEyOSAxMC4xODY0IDMzLjI0MzIgMTAuNTM4NCAzMi44NDA0QzExLjMyNTkgMzIuMTQ0MSAxMi4zNTMzIDMxLjc4NzcgMTMuMzk5MiAzMS44NDgxSDQ0LjIxNkM0NS4wNjg0IDI4LjU1NDQgNDYuMTgwOCAyNS4zMzQ4IDQ3LjU0MjUgMjIuMjE5OUgxMy4zOTkyQzExLjkwNzcgMjIuMjAxMSAxMC40MjIzIDIyLjQxMzcgOC45OTQ5MSAyMi44NTAyQzcuNzcxNDEgMjMuMjMyNyA2LjYxMzU2IDIzLjgwMjUgNS41NjE5NSAyNC41Mzk3QzQuNTk5NCAyNS4yMTYxIDMuNzM4OTEgMjYuMDI5MSAzLjAwNzE5IDI2Ljk1MzVDMi4zMjQ2NyAyNy44Mzg1IDEuNzQ4NDkgMjguODAxNyAxLjI5MDcgMjkuODIzMkMwLjg1NTUwMiAzMC43OTgyIDAuNTI5NzA5IDMxLjgxOTMgMC4zMTkzNjggMzIuODY3MkMwLjExNTMwNiAzMy44MDU2IDAuMDA4Mjc4NDkgMzQuNzYyOCAxLjU3NzdlLTA1IDM1LjcyMzRWNjEuODA1M0MtMC4wMTczNTg5IDYzLjMxNzIgMC4xOTM1MTUgNjQuODIyOCAwLjYyNTQwMSA2Ni4yNzA4QzEuMDA1MDggNjcuNTA3OSAxLjU3MDUgNjguNjc5MiAyLjMwMTk4IDY5Ljc0MzlDMi45NzAwNSA3MC43MTI1IDMuNzc3MjEgNzEuNTc1NiA0LjY5NzA1IDcyLjMwNTFDNS41Nzg1OCA3Mi45ODQyIDYuNTMzODggNzMuNTYgNy41NDQ1NCA3NC4wMjE2QzguNTE0MzIgNzQuNDU0MiA5LjUyNjc4IDc0Ljc4MjQgMTAuNTY1IDc1LjAwMDVDMTEuNDk2MiA3NS4yMDYyIDEyLjQ0NiA3NS4zMTQgMTMuMzk5MiA3NS4zMjIzSDQ3LjgyMTlDNDYuNDUyNCA3Mi4xOTcxIDQ1LjMyNjkgNjguOTY5IDQ0LjQ1NTUgNjUuNjY3NFoiIGZpbGw9IndoaXRlIi8+DQo8L2c+DQo8ZGVmcz4NCjxjbGlwUGF0aCBpZD0iY2xpcDBfODI2XzE0NTc2Ij4NCjxyZWN0IHdpZHRoPSIyMjEiIGhlaWdodD0iMTAzIiBmaWxsPSJ3aGl0ZSIvPg0KPC9jbGlwUGF0aD4NCjwvZGVmcz4NCjwvc3ZnPg0K';
export default class ThemeService {


  getServiceType(): ServiceTypeEnum {
    return ServiceTypeEnum.THEME_SERVICE;
  }

  getUserTheme(): ThemeInfo[] {
    //API call to get list of colours, for now just mockup
    let themeInfo: ThemeInfo[] = [];
    if(useUserStore()._isLoggedIn) {
      if(useUserStore().providerId == '77') {
        themeInfo = [
          {propertyName: '--main-color', propertyValue: '#3C3C3B'},
          {propertyName: '--secondary-color', propertyValue: '#95C11F'},
          {propertyName: '--main-lighter-color', propertyValue: '#535353'},
          {propertyName: '--main-darker-color', propertyValue: '#202020'},
          {propertyName: '--secondary-lighter-color', propertyValue: '#bbC11F'},
          {propertyName: '--secondary-darker-color', propertyValue: '#789b1a'},
          {propertyName: '--header-text-color', propertyValue: '#FFFFFF'},
          {propertyName: '--plain-text-color', propertyValue: '#152036'},
        ];

      } else {
        themeInfo = [
          {propertyName: '--main-color', propertyValue: '#002C50'},
          {propertyName: '--secondary-color', propertyValue: '#2AFD88'},
          {propertyName: '--main-lighter-color', propertyValue: '#02447A'},
          {propertyName: '--secondary-lighter-color', propertyValue: '#20DF76'},
          {propertyName: '--header-text-color', propertyValue: '#FFFFFF'},
          {propertyName: '--plain-text-color', propertyValue: '#152036'},
        ];
      }
    } else {
      themeInfo = [
        {propertyName: '--main-color', propertyValue: '#002C50'},
        {propertyName: '--secondary-color', propertyValue: '#2AFD88'},
        {propertyName: '--main-lighter-color', propertyValue: '#02447A'},
        {propertyName: '--secondary-lighter-color', propertyValue: '#20DF76'},
        {propertyName: '--header-text-color', propertyValue: '#FFFFFF'},
        {propertyName: '--plain-text-color', propertyValue: '#152036'},
      ];
    }
    return themeInfo;
  }

  getLogoURL() {
    if(useUserStore()._isLoggedIn) {
      useThemeStore().setLogoURL(logo);
    } else {
      useThemeStore().setLogoURL(logo);
    }
  }
}
