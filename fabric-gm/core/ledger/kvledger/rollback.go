/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kvledger

import (
	"github.com/VoneChain-CS/fabric-gm/common/ledger/blkstorage"
	"github.com/VoneChain-CS/fabric-gm/common/ledger/util/leveldbhelper"
	"github.com/pkg/errors"
)

// RollbackKVLedger rollbacks a ledger to a specified block number
func RollbackKVLedger(rootFSPath, ledgerID string, blockNum uint64) error {
	fileLockPath := fileLockPath(rootFSPath)
	fileLock := leveldbhelper.NewFileLock(fileLockPath)
	if err := fileLock.Lock(); err != nil {
		return errors.Wrap(err, "as another peer node command is executing,"+
			" wait for that command to complete its execution or terminate it before retrying")
	}
	defer fileLock.Unlock()

	blockstorePath := BlockStorePath(rootFSPath)
	if err := blkstorage.ValidateRollbackParams(blockstorePath, ledgerID, blockNum); err != nil {
		return err
	}

	logger.Infof("Dropping databases")
	if err := dropDBs(rootFSPath); err != nil {
		return err
	}

	logger.Info("Rolling back ledger store")
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}
	if err := blkstorage.Rollback(blockstorePath, ledgerID, blockNum, indexConfig); err != nil {
		return err
	}
	logger.Infof("The channel [%s] has been successfully rolled back to the block number [%d]", ledgerID, blockNum)
	return nil
}
