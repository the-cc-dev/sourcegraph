import { Observable, zip } from 'rxjs'
import { catchError, map, switchMap } from 'rxjs/operators'

import { ERPRIVATEREPOPUBLICSOURCEGRAPHCOM } from '../../../shared/backend/errors'
import { resolveRev, retryWhenCloneInProgressError } from '../../../shared/repo/backend'
import { FileInfo } from '../code_intelligence'

export const ensureRevisionsAreCloned = (files: Observable<FileInfo>): Observable<FileInfo> =>
    files.pipe(
        switchMap(({ repoName, commitID, baseCommitID, ...rest }) => {
            // Although we get the commit SHA's from elesewhere, we still need to
            // use `resolveRev` otherwise we can't guarantee Sourcegraph has the
            // revision cloned.

            // Head
            const resolvingHeadRev = resolveRev({ repoName, rev: commitID }).pipe(retryWhenCloneInProgressError())

            const requests = [resolvingHeadRev]

            // If theres a base, resolve it as well.
            if (baseCommitID) {
                const resolvingBaseRev = resolveRev({ repoName, rev: baseCommitID }).pipe(
                    retryWhenCloneInProgressError()
                )

                requests.push(resolvingBaseRev)
            }

            return zip(...requests).pipe(
                map(() => ({ repoName, commitID, baseCommitID, ...rest })),
                catchError(err => {
                    if (err.code === ERPRIVATEREPOPUBLICSOURCEGRAPHCOM) {
                        return [{ repoName, commitID, baseCommitID, ...rest }]
                    } else {
                        throw err
                    }
                })
            )
        })
    )
