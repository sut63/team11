/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntBookingEdges,
    EntBookingEdgesFromJSON,
    EntBookingEdgesFromJSONTyped,
    EntBookingEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBooking
 */
export interface EntBooking {
    /**
     * BOOKINGDATE holds the value of the "BOOKING_DATE" field.
     * @type {string}
     * @memberof EntBooking
     */
    bOOKINGDATE?: string;
    /**
     * BORROWITEM holds the value of the "BORROW_ITEM" field.
     * @type {number}
     * @memberof EntBooking
     */
    bORROWITEM?: number;
    /**
     * PHONENUMBER holds the value of the "PHONE_NUMBER" field.
     * @type {string}
     * @memberof EntBooking
     */
    pHONENUMBER?: string;
    /**
     * TIMELEFT holds the value of the "TIME_LEFT" field.
     * @type {string}
     * @memberof EntBooking
     */
    tIMELEFT?: string;
    /**
     * USERNUMBER holds the value of the "USER_NUMBER" field.
     * @type {number}
     * @memberof EntBooking
     */
    uSERNUMBER?: number;
    /**
     * 
     * @type {number}
     * @memberof EntBooking
     */
    clientID?: number;
    /**
     * 
     * @type {EntBookingEdges}
     * @memberof EntBooking
     */
    edges?: EntBookingEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBooking
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntBooking
     */
    servicepointID?: number;
    /**
     * 
     * @type {number}
     * @memberof EntBooking
     */
    userID?: number;
}

export function EntBookingFromJSON(json: any): EntBooking {
    return EntBookingFromJSONTyped(json, false);
}

export function EntBookingFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBooking {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bOOKINGDATE': !exists(json, 'BOOKING_DATE') ? undefined : json['BOOKING_DATE'],
        'bORROWITEM': !exists(json, 'BORROW_ITEM') ? undefined : json['BORROW_ITEM'],
        'pHONENUMBER': !exists(json, 'PHONE_NUMBER') ? undefined : json['PHONE_NUMBER'],
        'tIMELEFT': !exists(json, 'TIME_LEFT') ? undefined : json['TIME_LEFT'],
        'uSERNUMBER': !exists(json, 'USER_NUMBER') ? undefined : json['USER_NUMBER'],
        'clientID': !exists(json, 'client_ID') ? undefined : json['client_ID'],
        'edges': !exists(json, 'edges') ? undefined : EntBookingEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'servicepointID': !exists(json, 'servicepoint_ID') ? undefined : json['servicepoint_ID'],
        'userID': !exists(json, 'user_ID') ? undefined : json['user_ID'],
    };
}

export function EntBookingToJSON(value?: EntBooking | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'BOOKING_DATE': value.bOOKINGDATE,
        'BORROW_ITEM': value.bORROWITEM,
        'PHONE_NUMBER': value.pHONENUMBER,
        'TIME_LEFT': value.tIMELEFT,
        'USER_NUMBER': value.uSERNUMBER,
        'client_ID': value.clientID,
        'edges': EntBookingEdgesToJSON(value.edges),
        'id': value.id,
        'servicepoint_ID': value.servicepointID,
        'user_ID': value.userID,
    };
}


