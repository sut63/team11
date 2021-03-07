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
/**
 * 
 * @export
 * @interface ControllersSearch
 */
export interface ControllersSearch {
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    clientName?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    operator?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    phoneNumber?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    symbolClient?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    symbolID?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    symbolPhone?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    symbolUser?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersSearch
     */
    userID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersSearch
     */
    userName?: string;
}

export function ControllersSearchFromJSON(json: any): ControllersSearch {
    return ControllersSearchFromJSONTyped(json, false);
}

export function ControllersSearchFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersSearch {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'clientName': !exists(json, 'clientName') ? undefined : json['clientName'],
        'operator': !exists(json, 'operator') ? undefined : json['operator'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'symbolClient': !exists(json, 'symbolClient') ? undefined : json['symbolClient'],
        'symbolID': !exists(json, 'symbolID') ? undefined : json['symbolID'],
        'symbolPhone': !exists(json, 'symbolPhone') ? undefined : json['symbolPhone'],
        'symbolUser': !exists(json, 'symbolUser') ? undefined : json['symbolUser'],
        'userID': !exists(json, 'userID') ? undefined : json['userID'],
        'userName': !exists(json, 'userName') ? undefined : json['userName'],
    };
}

export function ControllersSearchToJSON(value?: ControllersSearch | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'clientName': value.clientName,
        'operator': value.operator,
        'phoneNumber': value.phoneNumber,
        'symbolClient': value.symbolClient,
        'symbolID': value.symbolID,
        'symbolPhone': value.symbolPhone,
        'symbolUser': value.symbolUser,
        'userID': value.userID,
        'userName': value.userName,
    };
}


