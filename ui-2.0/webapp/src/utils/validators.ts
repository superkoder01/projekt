function validateREGON9(regon: string): boolean {
  const reg = /^[0-9]{9}$/;
  if (!reg.test(regon)) {
    return false;
  } else {
    const digits = ("" + regon).split("");
    let checksum = (8 * parseInt(digits[0]) + 9 * parseInt(digits[1]) + 2 * parseInt(digits[2]) + 3 * parseInt(digits[3]) + 4 * parseInt(digits[4]) + 5 * parseInt(digits[5]) + 6 * parseInt(digits[6]) + 7 * parseInt(digits[7])) % 11;
    if (checksum == 10) {
      checksum = 0;
    }
    return (parseInt(digits[8]) == checksum);
  }
}

function validateREGON14(regon: string): boolean {
  const reg = /^[0-9]{14}$/;
  if (!reg.test(regon)) {
    return false;
  } else {
    const digits = ("" + regon).split("");
    let checksum = (2 * parseInt(digits[0]) + 4 * parseInt(digits[1]) + 8 * parseInt(digits[2]) + 5 * parseInt(digits[3]) + 0 * parseInt(digits[4]) + 9 * parseInt(digits[5]) + 7 * parseInt(digits[6]) + 3 * parseInt(digits[7]) + 6 * parseInt(digits[8]) + 1 * parseInt(digits[9]) + 2 * parseInt(digits[10]) + 4 * parseInt(digits[11]) + 8 * parseInt(digits[12])) % 11;
    if (checksum == 10) {
      checksum = 0;
    }
    return (parseInt(digits[13]) == checksum);
  }
}

export function validateREGON(regon: string): boolean {
  if (regon.length === 9) {
    return validateREGON9(regon);
  } else if (regon.length === 14) {
    return validateREGON14(regon);
  }
  return false;
}

export function validatePESEL(pesel: string): boolean {
  const reg = /^[0-9]{11}$/;
  if (reg.test(pesel) == false) {
    return false;
  } else {
    const digits = ("" + pesel).split("");
    const month = parseInt(pesel.substring(2, 4));
    const peselMonthRange = (month >= 21 && month <= 32) || (month >= 1 && month <= 12);
    if ((parseInt(pesel.substring(4, 6)) > 31) || !peselMonthRange) {
      return false;
    } else {
      let checksum = (1 * parseInt(digits[0]) + 3 * parseInt(digits[1]) + 7 * parseInt(digits[2]) + 9 * parseInt(digits[3]) + 1 * parseInt(digits[4]) + 3 * parseInt(digits[5]) + 7 * parseInt(digits[6]) + 9 * parseInt(digits[7]) + 1 * parseInt(digits[8]) + 3 * parseInt(digits[9])) % 10;
      if (checksum == 0) checksum = 10;
      checksum = 10 - checksum;
      return (parseInt(digits[10]) == checksum);
    }
  }
}

export function validateNIP(nip: string): boolean {
  const nipWithoutDashes = nip.replace(/-/g, "");
  const reg = /^[0-9]{10}$/;
  if (reg.test(nipWithoutDashes) == false) {
    return false;
  } else {
    const digits = ("" + nipWithoutDashes).split("");
    const checksum = (6 * parseInt(digits[0]) + 5 * parseInt(digits[1]) + 7 * parseInt(digits[2]) + 2 * parseInt(digits[3]) + 3 * parseInt(digits[4]) + 4 * parseInt(digits[5]) + 5 * parseInt(digits[6]) + 6 * parseInt(digits[7]) + 7 * parseInt(digits[8])) % 11;
    return (parseInt(digits[9]) == checksum);
  }
}



