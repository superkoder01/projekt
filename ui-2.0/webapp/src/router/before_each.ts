import {NavigationGuardNext, RouteLocationNormalized, useRouter} from 'vue-router';
import {LogLevel} from '@/services/logger/log-level';
import {ServiceTypeEnum} from '@/services/logger/service-type.enum';
import {useUserStore} from '@/store/user.store';
import {LoggerService} from '@/services/logger/logger.service';
import {RequirePropsInfo} from "@/router/RequirePropsInfo";
import {PermissionsService} from "@/services/permissions/permissions.service";
import {PagesEnum} from "@/services/permissions/pages-enum";
import {RoleEnum} from "@/services/permissions/role-enum";

export function createRouterBeforeEach(logger: LoggerService) {
  useRouter().beforeEach((to, from, next) => {
    logger.logToConsole(LogLevel.DEBUG, ServiceTypeEnum.ROUTER, 'go form:' + JSON.stringify(from.name) + ' to:' + JSON.stringify(to.name));
    const permissionService = new PermissionsService();
    if (to.name === undefined) {
      //in case of wrong url (url not found in router definition)
      logger.logToConsole(LogLevel.INFO, ServiceTypeEnum.ROUTER, 'redirecting to after login page');
      next(permissionService.getAfterLoginPage());
      return;
    }
    if (to.matched.some(record => record.meta.requiresAuth)) {
      if (useUserStore().isLoggedIn) {
        if (!checkPageType(logger, permissionService, to)) {
          next(permissionService.getAfterLoginPage());
          return;
        }
        checkIfPropsIsSetOrRedirect(logger, to, next);
        return;
      }
      logger.logToConsole(LogLevel.INFO, ServiceTypeEnum.ROUTER, 'user not logged in, redirecting to login page');
      next('/login');
    } else {
      logger.logToConsole(LogLevel.INFO, ServiceTypeEnum.ROUTER, 'requiresAuth not defined or FALSE. Route: ' + JSON.stringify(to.name));
      next();
    }
  });
}

function checkPageType(logger: LoggerService, permissionService: PermissionsService, to: RouteLocationNormalized): boolean {
  if (to.meta.pageType !== undefined) {
    const canView = permissionService.canView(to.meta.pageType as PagesEnum);
    if (!canView) {
      logger.logToConsole(LogLevel.INFO, ServiceTypeEnum.ROUTER, 'the pageType ' + PagesEnum[to.meta.pageType as keyof typeof PagesEnum] + ' is restricted for :' + RoleEnum[useUserStore().currentRole]);
      return false;
    }
  }
  return true;
}

function checkIfPropsIsSetOrRedirect(logger: LoggerService, to: RouteLocationNormalized, next: NavigationGuardNext) {
  if (to.meta.checkRequiredProps !== undefined) {
    const temp = to.meta.checkRequiredProps as RequirePropsInfo
    if (temp.requiredProps !== undefined && to.params[temp.requiredProps]) {
      next();
    } else {
      logger.logToConsole(LogLevel.INFO, ServiceTypeEnum.ROUTER, "required props '" + temp.requiredProps + "' not defined , redirecting to:" + temp.redirectTo);
      next(temp.redirectTo);
    }
  } else {
    next();
  }
}
