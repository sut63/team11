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
    EntCategoryEdges,
    EntCategoryEdgesFromJSON,
    EntCategoryEdgesFromJSONTyped,
    EntCategoryEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCategory
 */
export interface EntCategory {
    /**
     * CategoryName holds the value of the "CategoryName" field.
     * @type {string}
     * @memberof EntCategory
     */
    categoryName?: string;
    /**
     * 
     * @type {EntCategoryEdges}
     * @memberof EntCategory
     */
    edges?: EntCategoryEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCategory
     */
    id?: number;
}

export function EntCategoryFromJSON(json: any): EntCategory {
    return EntCategoryFromJSONTyped(json, false);
}

export function EntCategoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCategory {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'categoryName': !exists(json, 'CategoryName') ? undefined : json['CategoryName'],
        'edges': !exists(json, 'edges') ? undefined : EntCategoryEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCategoryToJSON(value?: EntCategory | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'CategoryName': value.categoryName,
        'edges': EntCategoryEdgesToJSON(value.edges),
        'id': value.id,
    };
}


