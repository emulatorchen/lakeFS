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
pub struct FindMergeBaseResult {
    /// The commit ID of the merge source
    #[serde(rename = "source_commit_id")]
    pub source_commit_id: String,
    /// The commit ID of the merge destination
    #[serde(rename = "destination_commit_id")]
    pub destination_commit_id: String,
    /// The commit ID of the merge base
    #[serde(rename = "base_commit_id")]
    pub base_commit_id: String,
}

impl FindMergeBaseResult {
    pub fn new(source_commit_id: String, destination_commit_id: String, base_commit_id: String) -> FindMergeBaseResult {
        FindMergeBaseResult {
            source_commit_id,
            destination_commit_id,
            base_commit_id,
        }
    }
}

