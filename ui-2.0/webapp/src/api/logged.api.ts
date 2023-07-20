import { LoggerService } from '@/services/logger/logger.service';
import { LogLevel } from '@/services/logger/log-level';
import {ApiTypeEnum} from "@/services/logger/api-type.enum";

export abstract class LoggedApi {

  static logger:LoggerService = new LoggerService;

  protected abstract getApiType(): ApiTypeEnum;

  protected logToConsole (logLevel: LogLevel, message: string, ...data: string[]) {
    if (data !== undefined && data.length > 0) {
      LoggedApi.logger.logToConsole(logLevel, this.getApiType(), message, ...data);
    } else {
      LoggedApi.logger.logToConsole(logLevel, this.getApiType(), message);
    }
  }
}
