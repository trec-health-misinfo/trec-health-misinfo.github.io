---
layout: default
---

# TREC Health Misinformation Track (2021)

## Track Introduction

Web search engines are frequently used to help people make decisions about health-related issues.  Unfortunately, the web is filled with misinformation regarding the efficacy of treatments for health issues.  Search users may not be able to discern correct from incorrect information, nor credible from non-credible sources.  As a result of finding misinformation deemed by the user to be useful to their decision making task, they can make incorrect decisions that waste money and put their health at risk.

The TREC Health Misinformation track fosters research on retrieval methods that promote reliable and correct information over misinformation for health-related decision making tasks.

### Track communication

Track annoucements are made via [google groups](https://groups.google.com/forum/#!forum/trec-health-misinformation-track).  

## 2021 Track Task

### AdHoc Web Retrieval

**Task Description:** Participants devise search technologies that promote credible and correct information over incorrect information, with the assumption that correct information can better lead people to make correct decisions.

Each topic concerns itself with a health issue and a treatment for that issue.  The topics represent a user who is looking for information that is useful for making a decision about whether about whether or not the treatment is helpful or unhelpful for treating the health issue.  Better search result rankings will place very useful documents that are credible and correct at the top of ranking, and will not return incorrect information.  In this search task, incorrect information is considered harmful and should not be returned.  

For each topic, we have chosen a 'stance' for the topic on whether the treatment is helpful or unhelpful for the health issue.  **We do not claim to be providing medical advice, and medical decisions should never be made based on the stance we have chosen.**  If a treatment is considered 'helpful', then correct documents will those construed to be supportive of the treatment and incorrect documents will be those that would disuade the searcher from the treatment.  Likewise, an 'unhelpful' treatment should return documents that disuade the searcher from using the treatment and should avoid returning documents that are supportive of using the treatment.  For each topic, we have included an `evidence` link for a webpage that the topic author used as the basis for the topic's stance.  

The primary ad-hoc task is to use the topic's `query` or `description` for a topic and return a ranking of documents without use of any of the other topic fields.  Manual runs may also make use of the other fields of the topic, e.g. the `stance` and `evidence` fields, but these runs will need to declare the use of this data to allow us to distinguish them from the primary task.

#### Document Collection

This year we wil be using the [**noclean** version of the C4 dataset](https://huggingface.co/datasets/allenai/c4) used by [Google to train their T5 model](https://www.tensorflow.org/datasets/catalog/c4). The collection is comprised of text extracts from the April 2019 snapshot of Common Crawl. The Collection contains \~ 1B English documents.

You can download the corpus on a Debian/Ubuntu machine using the following commands ([see HuggingFace for further information](https://huggingface.co/datasets/allenai/c4)).
```
sudo apt-get install git-lfs 
git lfs install
GIT_LFS_SKIP_SMUDGE=1 git clone https://huggingface.co/datasets/allenai/c4
cd c4
git lfs pull --include="en.noclean/c4-train*"
```
The collection is made up of the 7168 gzipped jsonl files located in the en.nolean directory.  We are using only the `c4-train.*.json.gz` files and not the `c4-validation.*.json.gz` files.  Each file containes \~150k documents, and has one document per line. A document is a json object with the fields `text`, `url` and `timestamp`. As packaged in c4.noclean, documents do not contain a document identifier. The document identifier ("docno") for each document will be `c4nc-<file#>-<line#>`, where `<file#>` is 4 digits and `<line#>` is 5 digits.  Both `<file#>` and `<line#>` are padded with zeros. We number lines (`<line#>`) in a file starting at 0. The `<file#>` is taken from from the file name. For example, the document on the second line of file c4-train.01234-of-07168.json.gz would have a docno of c4nc-1234-00001.

One way to insert document identifiers is by using the [provided script](renamer.go). Another would be name the documents as you index them.

#### Topics

To be released.

#### Evaluation

The final qrels will contain assessments with respect to the following criteria:
* *Usefulness*: The extent to which the document contains information that a search user would find useful in answering the topic's question.  Three levels: 0. Not useful. 1.  Useful: User would find the document useful for the topic because it provides useful information about the health issue, the treatment, or both. 2. Very useful: User finds the document very useful because it specifically talks about the use of the treatment for the health issue, or provides strong guidance about the health treatment regardless of the health issue.
* *Supportiveness*: Does the document support or dissuade the use of the treatment in the topic's question?  Three levels: 1. Supportive: The document would support a decision to use the treatment.  0. Neutral: The document neither supports or dissuades use of the treatment. -1. Dissuades: The document would dissuade a user to not use the treatment.
* *Credibility*: whether the document is considered credible by the assessor. 0. Not credible.  1. Credible.

*Note:* not-useful documents will not be assessed with respect to credibility and supportiveness. 

Submitted runs will be evaluated with respect to the three criteria: usefulness, correctness, and credibility. We will be using the [compatibility](https://github.com/claclark/Compatibility) measure in a similar fashion as we did in the 2020 version of this track (See [overview paper](https://trec.nist.gov/pubs/trec29/papers/OVERVIEW.HM.pdf) for more details). 

Details of the assessing instructions will be published after they are finalized. 

#### Runs 

Participating groups will be allowed to submit as many runs as they like, but they need authorization from the Track organizers before submitting more than 10 runs. Not all runs are likely to be used for pooling and groups will need to specify a preference ordering for pooling purposes.

Runs may be either automatic or manual. 

*Automatic runs:* Only the topic's `query` or `description` field may be used for automatic runs.  Use of the other fields, e.g. `stance`, `evidence`, and `narrative`, will make a run a manual run.  An automatic run is made without any tuning or customization based on the topics.  Best practice for an automatic run is to avoid using the topics or even looking at them until all decisions and code have been written to produce the automatic run. 

*Manual runs:* A manual run is anything that is not an automatic run. Manual runs commonly have some human input based on the topics, e.g., hand-crafted queries or relevance feedback. All topic fields may be used for manual runs.  We encourage manual runs in addition to automatic runs.  

Submission format for runs will follow the standard TREC run format.  For each topic, please return 1,000 ranked documents.  The standard TREC run format is as follows:

```
qid Q0 docno rank score tag
```
where:
* `qid`: the topic number;
* `Q0`: unused and should always be Q0;
* `docno`: the official document id number returned by your system for the topic `qid`;
* `ran`: the rank the document is retrieved;
* `score`: the score (integer or floating point) that generated the ranking. The score must be in descending (non-increasing) order. The score is important to handle tied scores. (`trec_eval` sorts documents by the specified scores values and not your ranks values);
* `tag`: a tag that uniquely identifies your group AND the method you used to produce the run. Each run should have a different tag.

The fields should be separated with a space. 

An example run is shown below:
```
1 Q0 c4nc-4124-069102 1 14.8928003311 myGroupNameMyMethodName
1 Q0 c4nc-3346-052165 2 14.7590999603 myGroupNameMyMethodName
1 Q0 c4nc-3904-054203 3 14.5707998276 myGroupNameMyMethodName
...
```

### Other Track Tasks

This year, we will only have the ad-hoc retrieval task.  We will consider adding back in the other tasks in future years.  

## Schedule
* **June 21, 2021** Collection released;
* **July 2021** Topics will be released;
* **September 2, 2021** Runs due;
* **End of September 2021** Results returned;
* **October 2021** Notebook paper due;
* **November 17-19, 2021** TREC Conference;
* **February 2022** Final report due.

## Organizers

* [Charles Clarke, University of Waterloo](https://cs.uwaterloo.ca/about/people/claclark)
* [Maria Maistro, University of Copenhagen](https://di.ku.dk/english/staff/?pure=en/persons/641366)
* [Mark Smucker, University of Waterloo](http://mansci.uwaterloo.ca/~msmucker/)

