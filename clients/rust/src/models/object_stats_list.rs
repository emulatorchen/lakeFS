/*
 * lakeFS API
 *
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: services@treeverse.io
 * Generated by: https://openapi-generator.tech
 */

use crate::models;

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ObjectStatsList {
    #[serde(rename = "pagination")]
    pub pagination: Box<models::Pagination>,
    #[serde(rename = "results")]
    pub results: Vec<models::ObjectStats>,
}

impl ObjectStatsList {
    pub fn new(pagination: models::Pagination, results: Vec<models::ObjectStats>) -> ObjectStatsList {
        ObjectStatsList {
            pagination: Box::new(pagination),
            results,
        }
    }
}

