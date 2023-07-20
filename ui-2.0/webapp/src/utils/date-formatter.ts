import moment from "moment";
import {DateFormat} from "@/models/billing/billing";

export function formatSendDate(dateString: string, outFormat: DateFormat): string {

  return moment(dateString).format(outFormat);
}

const onErrorStringDefault = "-";

export function formatDate(dateString: string, format: DateFormat, onErrorString?: string): string {
  console.log("dateString:" + dateString + " format:" + format + " onErrorString:" + onErrorString);
  if (dateString === undefined || dateString.length == 0) {
    if (onErrorString != undefined) {
      return onErrorString;
    } else {
      return onErrorStringDefault;
    }
  } else {
    const temp = moment(dateString).format(format);
    if (temp == "Invalid date") {
      if (onErrorString != undefined) {
        return onErrorString;
      } else {
        return onErrorStringDefault;
      }
    } else {
      return temp;
    }
  }
}


