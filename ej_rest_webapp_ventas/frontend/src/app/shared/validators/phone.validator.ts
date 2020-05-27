import {FormControl} from '@angular/forms';

export const validatePhone=(c: FormControl) => {
    if(!c.value) {
        return null;
    }

    const PHONE_REGEXP: any=/^[(]\d{3}[)]\d{3}[-]\d{4}$/;
    const result=PHONE_REGEXP.test(c.value)? null
        :{
            validatePhone: {
                valid: false
            }
        }
    return result;
}