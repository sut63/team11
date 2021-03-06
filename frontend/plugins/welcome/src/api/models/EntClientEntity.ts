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
    EntClientEntityEdges,
    EntClientEntityEdgesFromJSON,
    EntClientEntityEdgesFromJSONTyped,
    EntClientEntityEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClientEntity
 */
export interface EntClientEntity {
    /**
     * CLIENTNAME holds the value of the "CLIENT_NAME" field.
     * @type {string}
     * @memberof EntClientEntity
     */
    cLIENTNAME?: string;
    /**
     * 
     * @type {EntClientEntityEdges}
     * @memberof EntClientEntity
     */
    edges?: EntClientEntityEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntClientEntity
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntClientEntity
     */
    statusID?: number;
}

export function EntClientEntityFromJSON(json: any): EntClientEntity {
    return EntClientEntityFromJSONTyped(json, false);
}

export function EntClientEntityFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClientEntity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cLIENTNAME': !exists(json, 'CLIENT_NAME') ? undefined : json['CLIENT_NAME'],
        'edges': !exists(json, 'edges') ? undefined : EntClientEntityEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'statusID': !exists(json, 'status_ID') ? undefined : json['status_ID'],
    };
}

export function EntClientEntityToJSON(value?: EntClientEntity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'CLIENT_NAME': value.cLIENTNAME,
        'edges': EntClientEntityEdgesToJSON(value.edges),
        'id': value.id,
        'status_ID': value.statusID,
    };
}


